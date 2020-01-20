package _442_数组中重复的数据_find_all_duplicates_in_an_array

import "testing"

func TestFindDuplicates(t *testing.T) {
	ints := FindDuplicates([]int{10,2,5,10,9,1,1,4,3,7})
	t.Log(ints)
}
