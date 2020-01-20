package main

import "github.com/DestinyWang/go-widget/leetcode"

// 反转一个单链表

func ReverseList(node *leetcode.ListNode) *leetcode.ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	var p, q, t *leetcode.ListNode
	for p = node; p != nil; {
		q = p.Next
		p.Next = t
		t = p
		p = q
	}
	return t
}

func main() {

}
