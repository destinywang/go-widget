package _701_二叉搜索树中的插入操作_insert_into_a_binary_search_tree

import (
	"github.com/DestinyWang/go-widget/leetcode"
	"testing"
)

func TestInsertIntoBST(t *testing.T) {
	var root *leetcode.TreeNode
	root = InsertIntoBST(root, 5)
	root = InsertIntoBST(root, 3)
	root = InsertIntoBST(root, 10)
	root = InsertIntoBST(root, 1)
	root = InsertIntoBST(root, 2)
	t.Log(root)
}
