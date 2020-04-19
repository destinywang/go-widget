package _235_二叉搜索树的最近公共祖先_lowest_common_ancestor_of_a_binary_search_tree

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
// 对于有根树 T 的两个结点 p, q, 最近公共祖先表示为一个结点 x, 满足 x 是 p, q 的祖先且 x 的深度尽可能大(一个节点也可以是它自己的祖先)

// Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 仅仅是普通二叉树的方式, 二叉搜索树依据自身性质可以进一步优化
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	sp1 := make([]*TreeNode, 0, 0)
	sp2 := make([]*TreeNode, 0, 0)
	sq1 := make([]*TreeNode, 0, 0)
	sq2 := make([]*TreeNode, 0, 0)
	_, s1 := InOrderTravel(root, sp1, sp2, p)
	_, s2 := InOrderTravel(root, sq1, sq2, q)
	short, long := make([]*TreeNode, 0, 0), make([]*TreeNode, 0, 0)
	if len(s1) >= len(s2) {
		short = s2
		long = s1
	} else {
		short = s1
		long = s2
	}
	for i := range short {
		if short[i].Val != long[i].Val {
			return short[i-1]
		}
	}
	return short[len(short)-1]
}

func InOrderTravel(root *TreeNode, s1 []*TreeNode, s2 []*TreeNode, v *TreeNode) ([]*TreeNode, []*TreeNode) {
	if root == nil {
		return s1, s2
	}
	// 入栈
	s1 = append(s1, root)
	if root.Val == v.Val {
		for _, r := range s1 {
			s2 = append(s2, r)
		}
	}
	if root.Left != nil {
		s1, s2 = InOrderTravel(root.Left, s1, s2, v)
	}
	if root.Right != nil {
		s1, s2 = InOrderTravel(root.Right, s1, s2, v)
	}
	// 出栈
	return s1[:len(s1)-1], s2
}
