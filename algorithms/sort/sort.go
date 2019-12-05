package sort

import "fmt"

// 冒泡排序会把正确的序列排到末尾, 因此断尾不断头, 内层遍历需要从 0 开始
func BubbleSort(arr []int) {
	if len(arr) <= 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] <= arr[j+1] {
				t := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
			}
			fmt.Printf("i=[%d], j=[%d], arr=%v\n", i, j, arr)
		}
	}
}

func InsertionSort(arr []int) {
	if len(arr) <= 0 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				t := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = t
			}
		}
	}
}

// 选择排序的已排区间从 0 开始
func SelectionSort(arr []int) {
	if len(arr) <= 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] > arr[min] {
				min = j
			}
		}
		t := arr[i]
		arr[i] = arr[min]
		arr[min] = t
	}
}

func MergeSort(arr []int) {
	if len(arr) == 0 {
		return
	}

}

func mergeSort(arr []int, r int, l int) []int {
	if r >= l {
		return
	}
	mid := (r + l) / 2
	leftArr := mergeSort(arr, l, mid)
	rightArr := mergeSort(arr, mid+1, r)

}

func merge(leftArr []int, rightArr []int) []int {

}
