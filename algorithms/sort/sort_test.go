package sort

import (
	"github.com/Pallinder/go-randomdata"
	"testing"
)

var arr = []int{4, 6, 2, 1, 5, 3}

func TestBubbleSort(t *testing.T) {
	BubbleSort(arr)
	t.Log(arr)
}

func TestInsertionSort(t *testing.T) {
	InsertionSort(arr)
	t.Log(arr)
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(arr)
	t.Log(arr)
}

func TestMergeSort(t *testing.T) {
	MergeSort(arr)
	t.Log(arr)
}

func TestQuickSort(t *testing.T) {
	QuickSort(arr)
	t.Log(arr)
}

func TestHeapSort(t *testing.T) {
	arr = make([]int, 0)
	for i := 0; i < 50; i ++ {
		arr = append(arr, randomdata.Number(0, 100))
	}
	t.Log(HeapSort(arr))
}
