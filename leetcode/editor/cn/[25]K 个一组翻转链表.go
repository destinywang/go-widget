package cn

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。 
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。 
//
// 
//
// 示例： 
//
// 给你这个链表：1->2->3->4->5 
//
// 当 k = 2 时，应当返回: 2->1->4->3->5 
//
// 当 k = 3 时，应当返回: 3->2->1->4->5 
//
// 
//
// 说明： 
//
// 
// 你的算法只能使用常数的额外空间。 
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。 
// 
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var vals [][]int
	for p, i := head, 0; p != nil; p = p.Next {
		if i%k == 0 {
			vals = append(vals, []int{p.Val})
		} else {
			vals[i/k] = append(vals[i/k], p.Val)
		}
		i++
	}
	// insert to tail
	head = &ListNode{
		Val:  -1,
		Next: nil,
	}
	var tail = head
	for _, vs := range vals {
		if len(vs) == k {
			for i := len(vs) - 1; i >= 0; i-- {
				tail.Next = &ListNode{
					Val: vs[i],
				}
				tail = tail.Next
			}
		} else {
			for _, v := range vs {
				tail.Next = &ListNode{
					Val: v,
				}
				tail = tail.Next
			}
		}
	}
	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
