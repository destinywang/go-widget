package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{1, 2, 3, 5, 1, 2}
	t.Log(RemoveElement(nums, 1))
	t.Logf("nums=%v", nums)
}
