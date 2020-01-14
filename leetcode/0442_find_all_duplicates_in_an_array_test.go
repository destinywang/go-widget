package leetcode

import "testing"

func TestFindDuplicates(t *testing.T) {
	ints := FindDuplicates([]int{10,2,5,10,9,1,1,4,3,7})
	t.Log(ints)
}
