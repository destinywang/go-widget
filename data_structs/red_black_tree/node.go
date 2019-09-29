package red_black_tree

import (
	"fmt"
	"math"
)

// 1. 每个节点都是红色或黑色
// 2. 根节点必须为黑色
// 3. 叶子节点必须为黑色
// 4. 双亲节点为红色时, 叶子节点必须为黑色
// 5. 任意根所有路径的黑色节点数量相同
type Node struct {
	k                   int
	v                   interface{}
	color               Color
	parent, left, right *Node
}

type Color string

const Red Color = "RED"
const Black Color = "BLACK"

func (n *Node) Add(k int, v interface{}, parent *Node) *Node {
	if n == nil {
		color := Red
		if parent == nil {
			color = Black
		}
		return &Node{
			k:      k,
			v:      v,
			color:  color,
			parent: parent,
		}
	}
	if k == n.k {
		tmp := &Node{
			k:      n.k,
			v:      n.v,
			color:  n.color,
			parent: parent,
		}
		n.v = v
		return tmp
	}

	if k < n.k {
		n.left = n.left.Add(k, v, n)
	}
	if k > n.k {
		n.right = n.right.Add(k, v, n)
	}
	return n
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return int(math.Max(float64(n.left.Height()), float64(n.right.Height()))) + 1
}

func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.left.InOrder()
	fmt.Println(n.k)
	n.right.InOrder()
}
