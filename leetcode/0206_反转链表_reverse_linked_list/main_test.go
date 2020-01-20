package main

import (
	"github.com/DestinyWang/go-widget/leetcode"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := leetcode.InitLinkedList([]int{1, 2, 3, 4, 5})
	t.Log(list.String())
	reverseList := ReverseList(list)
	t.Log(reverseList.String())
}
