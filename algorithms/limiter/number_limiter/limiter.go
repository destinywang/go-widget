package number_limiter

import "time"

type Limiter interface {
	// 根据某个实体的特定行为在一定时间内的最大次数限制来判断当前行为能否继续
	// @param id: 实体 id, 如 userId 等
	// @param actionKey: 实体的某个特定行为, 如 "login"
	// @param duration: 参与统计的时间范围, 如一分钟内
	// @param maxCnt: 指定范围内允许发生的最大次数
	// @return 某个实体的特定行为在一定时间段内是否已经超过了最大允许次数
	IsActionAllowed(id string, actionKey string, duration time.Duration, maxCnt int) bool
}
