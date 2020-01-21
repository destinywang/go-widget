package main

import (
	"github.com/DestinyWang/go-widget/leetcode"
	"testing"
)

func TestTree2Str(t *testing.T) {
	t.Log(Tree2Str(leetcode.InitTreeNode([]int{1, 2, 3, 4, -1, -1, -1})))
	t.Log(Tree2Str(leetcode.InitTreeNode([]int{1, 2, 3, -1, 4, -1, -1})))
}
