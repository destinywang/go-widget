package _268_缺失数字_missing_number

// 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

func MissingNumber(nums []int) int {
	n := len(nums)
	cnt := 0
	for _, i := range nums {
		cnt += i
	}
	return n*(n+1)/2 - cnt
}
