package _203_移除链表元素_remove_linked_list_elements_test

import (
	"github.com/DestinyWang/go-widget/leetcode"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	head := &leetcode.ListNode{
		Val:1,
	}
	head = RemoveElements(head, 1)
	t.Log(head)
}
