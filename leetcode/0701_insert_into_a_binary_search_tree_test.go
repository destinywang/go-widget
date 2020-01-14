package leetcode

import "testing"

func TestInsertIntoBST(t *testing.T) {
	var root *TreeNode
	root = InsertIntoBST(root, 5)
	root = InsertIntoBST(root, 3)
	root = InsertIntoBST(root, 10)
	root = InsertIntoBST(root, 1)
	root = InsertIntoBST(root, 2)
	t.Log(root)
}
