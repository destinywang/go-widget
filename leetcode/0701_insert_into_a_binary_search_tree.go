package leetcode

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
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
