package _236_二叉树的最近公共祖先_lowest_common_ancestor_of_a_binary_tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val:   6,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	t.Log(LowestCommonAncestor(root, root.Left, root.Left.Left).Val)
}
