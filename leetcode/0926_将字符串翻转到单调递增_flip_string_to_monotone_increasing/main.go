package _926_将字符串翻转到单调递增_flip_string_to_monotone_increasing

//如果一个由 '0' 和 '1' 组成的字符串，是以一些 '0'（可能没有 '0'）后面跟着一些 '1'（也可能没有 '1'）的形式组成的，那么该字符串是单调递增的。
//我们给出一个由字符 '0' 和 '1' 组成的字符串 S，我们可以将任何 '0' 翻转为 '1' 或者将 '1' 翻转为 '0'。
//返回使 S 单调递增的最小翻转次数。

func MinFlipsMonoIncr(S string) int {
	dp := make([]int, 0, len(S))
	// init
	dp[0] = 0
	for i := 1; i < len(S); i++ {
	}
	return 0
}