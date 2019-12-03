package stack

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	stack.Push("e")
	stack.Push("f")
	t.Log(stack.Top())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
}
