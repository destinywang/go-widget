package huffman_tree

import (
	"fmt"
	"strings"
)

// 给定 n 个权值作为 n 个叶子节点, 构造一棵二叉树, 若该数的带权路径长度达到最小,
// 则称这样的树为 Huffman Tree, 权值越大离根节点越近
//            A
//           /\
//          /  \
//         /    \
//        B    C(10)
//       /\
//      /  \
//     /    \
//   D(13) E(15)
// 节点D的带权路径长度为 根到该节点的路径长度(2) * 节点的权(13) = 26
// 树的带权路径长度: 所有的叶子节点的带权路径长度之和, 即为 WPL(weighted path length)
// 上例的 WPL = 3*1 + 13*2 + 15*2 = 59
type TreeNode struct {
	value       int       // 权值
	data        string    // 字符
	left, right *TreeNode // 左右子树
}

func (tn *TreeNode) String() string {
	str := make([]string, 0)
	return strings.Join(preOrder(tn, str), "")
}

func preOrder(root *TreeNode, str []string) []string {
	if root == nil {
		return str
	}
	str = append(str, fmt.Sprintf("{%p, value=%d, left=%p, right=%p}\n", root, root.value, root.left, root.right))
	str = preOrder(root.left, str)
	str = preOrder(root.right, str)
	return str
}
