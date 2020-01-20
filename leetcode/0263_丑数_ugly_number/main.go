package main

// 编写一个程序判断给定的数是否为丑数。
// 丑数就是只包含质因数 2, 3, 5 的正整数。

func IsUgly(num int) bool {
	if num < 2 {
		return false
	}
	for num%5 == 0 {
		num /= 5
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%2 == 0 {
		num /= 2
	}
	return num == 1
}

func main() {

}
