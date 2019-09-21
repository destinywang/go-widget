package huffman_tree

import (
	"sort"
)

// 1. 对权值从小到大进行排序, 将每个权值节点都看做一棵二叉树
// 2. 取出根节点权值最小的两棵二叉树
// 3. 组成一棵新的二叉树, 新树的根节点是前面两棵树根节点权值之和
// 4. 将新树放入列表中重新排序, 重复 234 步直到列表中只剩下一棵二叉树
func CreateHuffmanTree(arr []int) *TreeNode {
	nodeList := make([]*TreeNode, 0, len(arr))
	for _, i := range arr {
		nodeList = append(nodeList, &TreeNode{
			value: i,
		})
	}
	for len(nodeList) > 1 {
		// step 1
		sort.Slice(nodeList, func(i, j int) bool {
			return nodeList[i].value < nodeList[j].value
		})
		// step 2
		leftNode := nodeList[0]
		rightNode := nodeList[1]
		// step 3
		parent := &TreeNode{
			value: leftNode.value + rightNode.value,
			left:  leftNode,
			right: rightNode,
		}
		nodeList = nodeList[2:]
		nodeList = append(nodeList, parent)
	}
	return nodeList[0]
}

// 编码方案: 向左路径为 0, 向右路径为 1
