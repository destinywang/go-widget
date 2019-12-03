package stack

type Stack struct {
	arr []interface{}
}

func NewStack() *Stack {
	return &Stack{
		arr: make([]interface{}, 0),
	}
}

func (stack *Stack) Push(data interface{}) {
	stack.arr = append(stack.arr, data)
}

func (stack *Stack) Pop() interface{} {
	if len(stack.arr) == 0 {
		return ""
	}
	data := stack.arr[len(stack.arr)-1]
	stack.arr = stack.arr[:len(stack.arr)-1]
	return data
}

func (stack *Stack) Top() interface{} {
	if len(stack.arr) == 0 {
		return ""
	}
	return stack.arr[len(stack.arr)-1]
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.arr) == 0
}
