package main

import "fmt"

// 给定两个字符串, A 和 B。
// A 的旋转操作就是将 A 最左边的字符移动到最右边。 例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。
// 如果在若干次旋转操作之后，A 能变成B，那么返回True。

func RotateString(A string, B string) bool {
	if A == B {
		return true
	}
	if len(A) != len(B) {
		return false
	}
	l := len(A)
	for i := range A {
		prefixA := A[:i+1]
		suffixA := A[i+1 : l]
		prefixB := B[:l-i-1]
		suffixB := B[l-i-1 : l]
		fmt.Printf("prefixA=[%s], suffixA=[%s], prefixB=[%s], suffixB=[%s]\n", prefixA, suffixA, prefixB, suffixB)
		if prefixA == suffixB && suffixA == prefixB {
			return true
		}
	}
	return false
}

func main() {

}
