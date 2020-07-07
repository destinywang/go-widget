package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	var vals []string
	for p := ln; p != nil; p = p.Next {
		vals = append(vals, strconv.FormatInt(int64(p.Val), 10))
	}
	return strings.Join(vals, "->")
}

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

func main() {
	var list = []int{1,2,3,4,5,6}
	fmt.Println(append(list[:3], list[4:]...))
}
