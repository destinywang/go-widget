package leetcode

import "math"

// 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
//找到所有出现两次的元素。
//你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

func FindDuplicates(nums []int) []int {
	repeats := make([]int, 0)
	for _, num := range nums {
		i := int(math.Abs(float64(num)) - 1)
		if nums[i] < 0 {
			repeats = append(repeats, i+1)
		} else {
			nums[i] *= -1
		}
	}
	return repeats
}
