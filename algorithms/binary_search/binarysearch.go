package binary_search

func BinarySearch(arr []int, value int) int {
	return binarySearch(arr, 0, len(arr), value)
}

func binarySearch(arr []int, l int, r int, value int) int {
	if len(arr) <= 0 || l >= r {
		return -1
	}
	for l < r {
		mid := l + (r-l)/2
		//fmt.Printf("l=[%d], r=[%d], mid=[%d]\n", l, r, mid)
		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}

// 查找第一个等于 value 的元素下标
func BinarySearch1(arr []int, value int) int {
	return binarySearch1(arr, 0, len(arr), value)
}

func binarySearch1(arr []int, l int, r int, value int) int {
	if len(arr) <= 0 || l >= r {
		return -1
	}
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == value {
			if mid == 0 || arr[mid-1] < value {
				return mid
			} else {
				r = mid - 1
			}
		} else if arr[mid] < value {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// 查询最后一个等于该值的元素下标
func BinarySearch2(arr []int, value int) int {
	return binarySearch2(arr, 0, len(arr), value)
}

func binarySearch2(arr []int, l int, r int, value int) int {
	if len(arr) <= 0 || l >= r {
		return -1
	}
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] == value {
			if mid == len(arr)-1 || arr[mid+1] > value {
				return mid
			}
			l = mid + 1
		} else if arr[mid] > value {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 查询第一个大于等于给定值的元素下标
func BinarySearch3(arr []int, value int) int {
	return binarySearch3(arr, 0, len(arr), value)
}

func binarySearch3(arr []int, l int, r int, value int) int {
	if len(arr) <= 0 || l >= r {
		return -1
	}
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] >= value {
			if mid == 0 || arr[mid-1] < value {
				return mid
			}
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 查询最后一个小于等于给定值的元素下标
func BinarySearch4(arr []int, value int) int {
	return binarySearch4(arr, 0, len(arr), value)
}

func binarySearch4(arr []int, l int, r int, value int) int {
	if len(arr) <= 0 || l >= r {
		return -1
	}
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] == value {
			if mid == len(arr)-1 || arr[mid+1] > value {
				return mid
			}
			l = mid + 1
		} else if arr[mid] < value {
			if arr[mid+1] > value {
				return mid
			}
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}
