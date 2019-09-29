package red_black_tree

import "testing"

func TestTree_Add(t *testing.T) {
	tree := new(Tree)
	tree.Add(10, 10)
	tree.Add(5, 5)
	tree.Add(15, 15)
	tree.Add(13, 13)
	tree.root.InOrder()
	t.Log(tree.root.Height())
}
