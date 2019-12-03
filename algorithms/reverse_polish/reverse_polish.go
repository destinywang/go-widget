package reverse_polish

import (
	"github.com/DestinyWang/go-widget/data_structs/stack"
	"strconv"
)

var numSet = map[uint8]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

var operMap = map[uint8]string{
	'+': "+", '-': "-", '*': "*", '/': "/", '(': "(", ')': ")",
}

var operSet = map[string]struct{}{
	"+": {}, "-": {}, "*": {}, "/": {}, "(": {}, ")": {},
}

func RePolish(expr string) (float64, error) {
	var (
		postfixExpr []string
		numStack    = stack.NewStack()
	)
	postfixExpr = resolve(expr)
	for _, s := range postfixExpr {
		if _, ok := operSet[s]; ok {
			// 如果是操作符, 从数值栈中弹出两个数字
			arg2 := numStack.Pop().(float64)
			arg1 := numStack.Pop().(float64)
			numStack.Push(cal(arg1, arg2, s))
		} else {
			float, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return 0, err
			}
			numStack.Push(float)
		}
	}
	return numStack.Pop().(float64), nil
}

// 将普通表达式转换为后缀表达式
func resolve(expr string) (subExpr []string) {
	var (
		p, q      int
		operStack = stack.NewStack()
	)
	q = 0
	for p = 0; p < len(expr); p++ {
		if _, ok := numSet[expr[p]]; ok {
			continue
		} else {
			subExpr = append(subExpr, expr[q:p])
			q = p + 1
		}
		// 如果是操作符
		if oper, ok := operMap[expr[p]]; ok {
			if operStack.IsEmpty() {
				operStack.Push(oper)
			} else {
				if priority(oper) < priority(operStack.Top().(string)) {
					// 如果当前操作符优先级低于栈顶
					for !operStack.IsEmpty() {
						subExpr = append(subExpr, operStack.Pop().(string))
					}
				}
				operStack.Push(oper)
			}
		}
	}
	if _, err := strconv.ParseInt(expr[q:p], 10, 64); err == nil {
		subExpr = append(subExpr, expr[q:p])
	}
	for !operStack.IsEmpty() {
		subExpr = append(subExpr, operStack.Pop().(string))
	}
	return subExpr
}

func priority(operation string) int {
	switch operation {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "(":
		return 10
	default:
		return 0
	}
}

func cal(arg1 float64, arg2 float64, oper string) float64 {
	switch oper {
	case "+":
		return arg1 + arg2
	case "-":
		return arg1 - arg2
	case "*":
		return arg1 * arg2
	case "/":
		return arg1 / arg2
	}
	return 0
}
