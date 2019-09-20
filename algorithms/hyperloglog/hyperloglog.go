package hyperloglog

type HyperLogLog interface {
	// 添加元素
	Add(key string, value string) bool
	// 统计总数
	Count(key string) int
}
