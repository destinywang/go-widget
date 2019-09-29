package red_black_tree

type Tree struct {
	root *Node
}

// 1. 红色是默认插入颜色(因为违规修复代价较小)
// 2. 如果没有双亲, 变为黑色
// 3. 双亲为黑色
// 4. 双亲为红色, 双亲需要变为黑色

func (t *Tree) Add(k int, v interface{}) interface{} {
	t.root = t.root.Add(k, v, nil)
	return t.root.v
}
