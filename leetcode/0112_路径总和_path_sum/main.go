package main

import (
	"fmt"
	"github.com/DestinyWang/go-widget/leetcode"
)

// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

func HasPathSum(root *leetcode.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	var flag bool
	hasPathSum0(root, sum, []int{}, &flag)
	return flag
}

func hasPathSum0(root *leetcode.TreeNode, sum int, stack []int, flag *bool) {
	stack = append(stack, root.Val)
	fmt.Printf("val=[%d], stack=%+v\n", root.Val, stack)
	var n int
	for _, i := range stack {
		n += i
	}
	if root.Left == nil && root.Right == nil && n == sum {
		*flag = true
	}
	if root.Left != nil {
		hasPathSum0(root.Left, sum, stack, flag)
	}
	if root.Right != nil {
		hasPathSum0(root.Right, sum, stack, flag)
	}
	stack = stack[:len(stack)-1]
}

func main() {

}
