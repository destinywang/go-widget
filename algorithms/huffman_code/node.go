package huffman_code

import (
	"fmt"
	"strings"
)

type Node struct {
	data        byte // 存放数据
	weight      int  // 权值
	left, right *Node
}

func (n *Node) String() string {
	strs := make([]string, 0)
	return strings.Join(preOrder(n, strs), "")
}

func preOrder(root *Node, strs []string) []string {
	if root == nil {
		return strs
	}
	strs = append(strs, fmt.Sprintf("%p, data=[%d], weight=[%d], left=[%p], right=[%p]\n", root, root.data, root.weight, root.left, root.right))
	strs = preOrder(root.left, strs)
	strs = preOrder(root.right, strs)
	return strs
}
