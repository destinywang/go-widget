package leetcode

import "strconv"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	for p := l; p != nil; p = p.Next {
		s = s + strconv.FormatInt(int64(p.Val), 10) + "->"
	}
	return s
}

func InitTreeNode(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(nums), len(nums))
	for i, num := range nums {
		if num == -1 {
			continue
		}
		nodes[i] = &TreeNode{
			Val: num,
		}
		p := (i - 1) / 2
		if p < 0 {
			continue
		}
		if i%2 == 1 {
			nodes[p].Left = nodes[i]
		} else {
			nodes[p].Right = nodes[i]
		}
	}
	return nodes[0]
}

func InitLinkedList(nums []int) *ListNode {
	if len(nums) < 1 {
		return nil
	}
	nodes := make([]*ListNode, len(nums), len(nums))
	for i := range nums {
		nodes[i] = &ListNode{
			Val: nums[i],
		}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	return nodes[0]
}
