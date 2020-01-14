package leetcode

// 删除链表中等于给定值 val 的所有节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
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
