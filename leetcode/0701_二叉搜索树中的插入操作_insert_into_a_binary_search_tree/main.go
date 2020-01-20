package _701_二叉搜索树中的插入操作_insert_into_a_binary_search_tree

import "github.com/DestinyWang/go-widget/leetcode"

func InsertIntoBST(root *leetcode.TreeNode, val int) *leetcode.TreeNode {
	if root == nil {
		return &leetcode.TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if val == root.Val {
		return root
	} else if val > root.Val {
		root.Right = InsertIntoBST(root.Right, val)
	} else {
		root.Left = InsertIntoBST(root.Left, val)
	}
	return root
}
