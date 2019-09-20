package throttle_limiter

type funnel interface {
	// 漏斗限流
	// @param key: 行为
	// @param capacity: 漏斗的初始容量
	// @param opsPerSecond: 每秒生成的令牌数量
	// @param quota: 每个行为占用的令牌数
	Throttle(key string, capacity int, opsPerSecond int, quota int)
}
