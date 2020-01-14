package leetcode

import "testing"

func TestRemoveElements(t *testing.T) {
	head := &ListNode{
		Val:1,
	}
	head = RemoveElements(head, 1)
	t.Log(head)
}
