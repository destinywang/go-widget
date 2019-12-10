package binary_search

import "testing"

var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
var arr1 = []int{1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 20}

func TestBinarySearch(t *testing.T) {
	t.Log(BinarySearch(arr, 10))
	t.Log(BinarySearch(arr, 1))
	t.Log(BinarySearch(arr, 20))
	t.Log(BinarySearch(arr, 0))
}

func TestBinarySearch1(t *testing.T) {
	t.Log(BinarySearch1(arr1, 1))
	t.Log(BinarySearch1(arr1, 2))
	t.Log(BinarySearch1(arr1, 3))
	t.Log(BinarySearch1(arr1, 20))
}

func TestBinarySearch2(t *testing.T) {
	t.Log(BinarySearch2(arr1, 1))
	t.Log(BinarySearch2(arr1, 2))
	t.Log(BinarySearch2(arr1, 3))
	t.Log(BinarySearch2(arr1, 20))
}

func TestBinarySearch3(t *testing.T) {
	t.Log(BinarySearch3(arr1, 1))
	t.Log(BinarySearch3(arr1, 2))
	t.Log(BinarySearch3(arr1, 3))
	t.Log(BinarySearch3(arr1, 20))
}

func TestBinarySearch4(t *testing.T) {
	t.Log(BinarySearch4(arr1, 1))
	t.Log(BinarySearch4(arr1, 2))
	t.Log(BinarySearch4(arr1, 3))
	t.Log(BinarySearch4(arr1, 20))
}
