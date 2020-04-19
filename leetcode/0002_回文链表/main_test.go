package _002_回文链表

import "testing"

func TestIsPalindrome(t *testing.T) {
	t.Log(IsPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}))
}
