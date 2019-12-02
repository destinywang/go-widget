package linked_list

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (linkedList *LinkedList) Pos(v string) int {
	pos := -1
	p := linkedList.Head
	if p == nil {
		return -1
	}
	for p.Next != nil {
		pos++
		if p.Data == v {
			return pos
		}
		p = p.Next
	}
	return -1
}

// 在指定位置后插入
func (linkedList *LinkedList) Insert(pos int, data string) bool {
	if linkedList.Head == nil {
		return false
	}
	p := linkedList.Head
	for i := 0; i < pos; i++ {
		p = p.Next
		if p == nil {
			return false
		}
	}
	t := &Node{
		Data: data,
		Next: p.Next,
	}
	p.Next = t
	if t.Next == nil {
		linkedList.Tail = t
	}
	return true
}

// 删除指定位置元素
func (linkedList *LinkedList) Del(pos int) bool {
	if linkedList.Head == nil {
		return false
	}
	p := linkedList.Head
	q := p
	for i := 0; i < pos-1; i++ {
		if i > 0 {
			q = q.Next
		}
		p = p.Next
		if p == nil {
			return false
		}
	}
	if p.Next == nil {
		q.Next = nil
		linkedList.Tail = q
	} else {
		q.Next = p.Next
	}
	return true
}

// 向链表尾部追加节点
func (linkedList *LinkedList) Append(v string) {
	if linkedList.Head == nil {
		linkedList.Head = &Node{
			Data: v,
			Next: nil,
		}
		linkedList.Tail = linkedList.Head
	} else {
		p := linkedList.Head
		for p = linkedList.Head; p.Next != nil; p = p.Next {
		}
		p.Next = &Node{
			Data: v,
			Next: nil,
		}
		linkedList.Tail = p.Next
	}
}

// 向链表头部插入节点
func (linkedList *LinkedList) Push(data string) {
	if linkedList.Head == nil {
		linkedList.Head = &Node{
			Data: data,
			Next: nil,
		}
		linkedList.Tail = linkedList.Head
	} else {
		newHead := &Node{
			Data: data,
			Next: linkedList.Head,
		}
		linkedList.Head = newHead
	}
}

func (linkedList *LinkedList) Reverse() {
	var (
		p, q, t *Node
	)
	if linkedList.Head == nil {
		return
	}
	p = linkedList.Head
	q = p.Next
	for q != nil {
		t = q.Next
		if p == linkedList.Head {
			p.Next = nil
		}
		q.Next = p
		p = q
		q = t
	}
	linkedList.Head = p
	for p = linkedList.Head; p.Next != nil; p = p.Next {
	}
	linkedList.Tail = p
}

func (linkedList *LinkedList) String() string {
	var strs []string
	p := linkedList.Head
	for p != nil {
		strs = append(strs, p.String())
		p = p.Next
	}
	return strings.Join(strs, " -> ") + fmt.Sprintf("head=[%s], tail=[%s]", linkedList.Head.Data, linkedList.Tail.Data)
}

func (linkedList *LinkedList) Pop() string {
	var p, q *Node
	q = linkedList.Head
	for p = linkedList.Head; p.Next != nil; p = p.Next {
		if p != linkedList.Head {
			q = q.Next
		}
	}
	q.Next = nil
	linkedList.Tail = q
	return p.Data
}
