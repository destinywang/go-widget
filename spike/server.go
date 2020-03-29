package main

import (
	"github.com/DestinyWang/go-widget/spike/util"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	redisPool       *redis.Pool
	done            chan int
	localSpike      LocalSpike
	remoteSpikeKeys RemoteSpikeKeys
)

func main() {
	http.HandleFunc("/ticket/buy", handleReq)
	http.ListenAndServe(":3005", nil)
}

func init() {
	localSpike = LocalSpike{
		LocalInStock:     150,
		LocalSalesVolume: 0,
	}
	remoteSpikeKeys = RemoteSpikeKeys{
		SpikeOrderHashKey:  "ticket_hash_key",
		TotalInventoryKey:  "ticket_total_nums",
		QuantityOfOrderKey: "ticket_sold_nums",
	}
	redisPool = NewPool()
	done = make(chan int, 1)
	done <- 1
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	var (
		logMg string
	)
	redisConn := redisPool.Get()
	<-done
	// 全局读写锁
	if localSpike.LocalDeductionStock() && remoteSpikeKeys.RemoteDeductionStock(redisConn) {
		util.RespJson(w, 1, "抢票成功", nil)
		logMg = logMg + "result:1,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
		// 发送消息, 创建订单
	} else {
		util.RespJson(w, -1, "已售罄", nil)
		logMg = logMg + "result:0,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
	}
	done <- 1
	//将抢票状态写入到log中
	writeLog(logMg, "./stat.log")
}

func writeLog(msg string, logPath string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	content := strings.Join([]string{msg, "\r\n"}, "")
	buf := []byte(content)
	fd.Write(buf)
}
