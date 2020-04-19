package _002_回文链表

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->2->1 true
// 1->2->2->2 false
func IsPalindrome(head *ListNode) bool {
	var s = make([]int, 0, 0)
	for p := head; p != nil; p = p.Next {
		s = append(s, p.Val)
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
