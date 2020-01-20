package main

import "math"

// 给定一个正整数 n，你可以做如下操作：
// 1. 如果 n 是偶数，则用 n / 2替换 n。
// 2. 如果 n 是奇数，则可以用 n + 1或n - 1替换 n。
// n 变为 1 所需的最小替换次数是多少？

func IntegerReplacement(n int) int {
	return integerReplacement0(n, 0)
}

func integerReplacement0(n, cnt int) int {
	if n == 1 {
		return cnt
	}
	if n%2 == 0 {
		return integerReplacement0(n/2, cnt+1)
	} else {
		return int(math.Min(float64(integerReplacement0(n+1, cnt+1)), float64(integerReplacement0(n-1, cnt+1))))
	}
}

func main() {

}
