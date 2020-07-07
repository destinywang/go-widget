package cn

//使用队列实现栈的下列操作：
//
// 
// push(x) -- 元素 x 入栈 
// pop() -- 移除栈顶元素 
// top() -- 获取栈顶元素 
// empty() -- 返回栈是否为空 
// 
//
// 注意: 
//
// 
// 你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合
//法的。 
// 你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。 
// 你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。 
// 
// Related Topics 栈 设计

//leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	l1 []int
	l2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if len(this.l1) != 0 {
		this.l1 = append(this.l1, x)
	} else {
		this.l2 = append(this.l2, x)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.l1) != 0 {
		for len(this.l1) > 1 {
			this.l2 = append(this.l2, this.l1[0])
			this.l1 = this.l1[1:]
		}
		n := this.l1[0]
		this.l1 = make([]int, 0, 0)
		return n
	} else {
		for len(this.l2) > 1 {
			this.l1 = append(this.l1, this.l2[0])
			this.l2 = this.l2[1:]
		}
		n := this.l2[0]
		this.l2 = make([]int, 0, 0)
		return n
	}
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.l1) != 0 {
		return this.l1[len(this.l1)-1]
	} else {
		return this.l2[len(this.l2)-1]
	}
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.l1) == 0 && len(this.l2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
