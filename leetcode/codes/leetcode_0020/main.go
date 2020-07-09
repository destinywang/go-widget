package main

import "fmt"

func isValid(s string) bool {
	var pattMap = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	for _, r := range s {
		if _, ok := pattMap[r]; !ok {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			if pattMap[r] != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
