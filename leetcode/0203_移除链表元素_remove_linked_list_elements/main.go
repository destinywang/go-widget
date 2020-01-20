package _203_移除链表元素_remove_linked_list_elements

import "github.com/DestinyWang/go-widget/leetcode"

// 删除链表中等于给定值 val 的所有节点

func RemoveElements(head *leetcode.ListNode, val int) *leetcode.ListNode {
	for ; head != nil && head.Val == val; head = head.Next {

	}
	if head == nil {
		return head
	}
	for p := head; p != nil && p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

// 链表在遍历删除时, 如果删除一个后继节点, 不能直接将 p 后移
