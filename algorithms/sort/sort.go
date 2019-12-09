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

// 左闭右闭, 只有 mid+1 才能推动递归分解, 所有元素都需要处理因此选择左闭右闭
func MergeSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (r + l) / 2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, m, r int) {
	fmt.Printf("l=[%d], m=[%d], r=[%d], subArr=%v\n", l, m, r, arr[l:r+1])
	var newArr []int
	p := l
	q := m + 1
	for p < m+1 && q <= r {
		if arr[p] < arr[q] {
			newArr = append(newArr, arr[p])
			p++
		} else {
			newArr = append(newArr, arr[q])
			q++
		}
	}
	var start, end int
	if p < m+1 {
		start = p
		end = m + 1
	} else {
		start = q
		end = r
	}
	for start < end {
		newArr = append(newArr, arr[start])
		start++
	}
	fmt.Println(newArr)
	for i := 0; i < len(newArr); i++ {
		arr[l+i] = newArr[i]
	}
}

//QuickSort 左闭右开
func QuickSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	pivot := arr[r]
	p := l
	fmt.Printf("arr=%v, l=[%d], r=[%d]\n", arr, l, r)
	for q := l; q < r; q++ {
		if arr[q] < pivot {
			t := arr[p]
			arr[p] = arr[q]
			arr[q] = t
			p++
		}
	}
	// swap p & r-1
	t := arr[p]
	arr[p] = arr[r]
	arr[r] = t
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}
