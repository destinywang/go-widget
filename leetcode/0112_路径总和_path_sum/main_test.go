package main

import (
	"github.com/DestinyWang/go-widget/leetcode"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	root := &leetcode.TreeNode{
		Val: 5,
		Left: &leetcode.TreeNode{
			Val: 4,
			Left: &leetcode.TreeNode{
				Val: 11,
				Left: &leetcode.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &leetcode.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &leetcode.TreeNode{
			Val: 8,
			Left: &leetcode.TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &leetcode.TreeNode{
				Val:  4,
				Left: nil,
				Right: &leetcode.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	t.Log(HasPathSum(root, 22))
	t.Log(HasPathSum(leetcode.InitTreeNode([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, -1, 1}), 22))
}
