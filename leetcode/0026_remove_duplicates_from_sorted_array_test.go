package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	ints := []int{1, 1, 2, 2, 3}
	length := RemoveDuplicates(ints)
	t.Logf("ints=%v", ints)
	t.Logf("length=%d", length)
}
