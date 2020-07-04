package cn

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 
//
// 
//
// 示例: 
//
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
// 
// Related Topics 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var fHead = &ListNode{
		Val:  0,
		Next: head,
	}
	var p, q, t *ListNode
	for t = fHead; t.Next != nil && t.Next.Next != nil; t = t.Next.Next {
		p, q = t.Next, t.Next.Next
		t.Next = q
		p.Next = q.Next
		q.Next = p
	}
	return fHead.Next
}

//leetcode submit region end(Prohibit modification and deletion)
