package _215_数组中的第k个最大元素_Kth_Largest_Element_in_an_Array

import (
	"github.com/DestinyWang/go-widget/data_structs/heap"
)

func FindKthLargest(nums []int, k int) int {
	return findByHeap(nums, k)
}

func findByHeap(nums []int, k int) int {
	// init heap
	h := heap.CreateHeap(nums)
	num := 0
	for i := 0; i < k; i++ {
		num = h.Pop()
	}
	return num
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
