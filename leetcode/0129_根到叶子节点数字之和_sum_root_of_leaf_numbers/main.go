package _129_根到叶子节点数字之和_sum_root_of_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNumbers(root *TreeNode) int {
	var stack []int
	sum := 0
	sumNumbers0(root, stack, &sum)
	return sum
}

func sumNumbers0(root *TreeNode, stack []int, sum *int) {
	stack = append(stack, root.Val)
	if root.Left == nil && root.Right == nil {
		*sum = *sum + Stack2Int(stack)
	}
	if root.Left != nil {
		sumNumbers0(root.Left, stack, sum)
	}
	if root.Right != nil {
		sumNumbers0(root.Right, stack, sum)
	}
}

func Stack2Int(stack []int) int {
	n := 0
	for i := 0; i < len(stack); i++ {
		n = n*10 + stack[i]
	}
	return n
}
