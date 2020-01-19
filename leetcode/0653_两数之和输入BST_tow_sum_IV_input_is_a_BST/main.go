package main

// 给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindTarget(root *TreeNode, k int) bool {
	val := make(map[int]struct{})
	return findTarget0(root, k, val)
}

func findTarget0(root *TreeNode, k int, val map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := val[root.Val]; ok {
		return true
	} else {
		val[k-root.Val] = struct{}{}
	}
	return findTarget0(root.Left, k, val) || findTarget0(root.Right, k, val)
}

func main() {

}
