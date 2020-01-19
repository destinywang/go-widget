package main

import "testing"

func TestSumNumbers(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	t.Log(SumNumbers(root))
	root = &TreeNode{
		Val:   1,
		Left:  &TreeNode{0, nil, nil},
		Right: nil,
	}
	t.Log(SumNumbers(root))
}
