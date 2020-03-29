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