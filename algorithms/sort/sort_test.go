package sort

import "testing"

var arr = []int{4, 6, 2, 3, 5, 1}

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
