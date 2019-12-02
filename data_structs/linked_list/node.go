package linked_list

import "fmt"

type Node struct {
	Data string
	Next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("[Addr=%p, Data=%s, Next=%p]\n", n, n.Data, n.Next)
}
