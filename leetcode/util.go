package leetcode

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
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
