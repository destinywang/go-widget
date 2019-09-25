package avl

type Tree struct {
	root *Node
}

func (at *Tree) Height() int {
	return at.root.Height()
}

func (at *Tree) Add(k, v int) {
	at.root = at.root.Add(k, v)
}

func (at *Tree) InOrder() {
	at.root.InOrder()
}

func (at *Tree) LeftHeight() int {
	return at.root.left.Height()
}

func (at *Tree) RightHeight() int {
	return at.root.right.Height()
}
