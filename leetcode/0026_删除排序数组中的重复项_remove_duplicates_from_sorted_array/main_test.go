package _026_删除排序数组中的重复项_remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	ints := []int{1, 1, 2, 2, 3}
	length := RemoveDuplicates(ints)
	t.Logf("ints=%v", ints)
	t.Logf("length=%d", length)
}
