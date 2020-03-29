package main

import (
	"github.com/gomodule/redigo/redis"
)

const LuaScript string = `
local ticket_key = KEYS[1]
local ticket_total_key = ARGV[1]
local ticket_sold_key = ARGV[2]
local ticket_total_nums = tonumber(redis.call('HGET', ticket_key, ticket_total_key))
local ticket_sold_nums = tonumber(redis.call('HGET', ticket_key, ticket_sold_key))
-- 查看是否有余票, 增加订单数量, 返回结果值
if (ticket_sold_nums > ticket_total_nums) then
    return redis.call('HINCRBY', ticket_key, ticket_sold_key, 1)
end
return 0
`

// 远程订单存储 key
type RemoteSpikeKeys struct {
	SpikeOrderHashKey  string // redis 秒杀订单 hash 结构 key
	TotalInventoryKey  string // hash 结构中总订单库存 key
	QuantityOfOrderKey string // hash 结构中已有订单数量 key
}

// 远端统一扣库存
func (remoteSpikeKeys *RemoteSpikeKeys) RemoteDeductionStock(conn redis.Conn) bool {
	lua := redis.NewScript(1, LuaScript)
	if result, err := redis.Int(lua.Do(conn, remoteSpikeKeys.SpikeOrderHashKey, remoteSpikeKeys.TotalInventoryKey, remoteSpikeKeys.QuantityOfOrderKey)); err != nil {
		return false
	} else {
		return result != 0
	}
}

// 初始化 redis 连接池
func NewPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			if conn, err = redis.Dial("tcp", ":6379"); err != nil {
				panic(err.Error())
			}
			return
		},
		MaxIdle:   10000,
		MaxActive: 12000,
	}
}
