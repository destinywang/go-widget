package main

func backspaceCompare(S string, T string) bool {
	var s1, s2 = InitStack(S), InitStack(T)
	var flag = true
	if len(s1) != len(s2) {
		flag = false
	} else {
		for i := range s1 {
			if s1[i] != s2[i] {
				flag = false
				break
			}
		}
	}
	return flag
}

func InitStack(S string) (stack []rune) {
	for _, r := range S {
		if len(stack) > 0 && r == '#' {
			stack = stack[:len(stack)-1]
		}
		if r != '#' {
			stack = append(stack, r)
		}
	}
	return stack
}
