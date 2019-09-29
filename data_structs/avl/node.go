package avl

import (
	"fmt"
	"math"
)

type Node struct {
	k           int
	v           interface{}
	left, right *Node
}

// 返回该节点的树的高度
func (t *Node) Height() int {
	if t == nil {
		return 0
	}
	if t.left == nil {
		return int(t.right.Height()) + 1
	}
	if t.right == nil {
		return int(t.left.Height()) + 1
	}
	return int(math.Max(float64(t.left.Height()), float64(t.right.Height()))) + 1
}

func (t *Node) Search(k interface{}) (v interface{}) {
	if t.k == k {
		return v
	}
	if t.left != nil {
		return t.left.Search(k)
	}
	if t.right != nil {
		return t.right.Search(k)
	}
	return nil
}

func (t *Node) Add(k int, v interface{}) *Node {
	if t == nil {
		return &Node{
			k: k,
			v: v,
		}
	} else {
		if t.k > k {
			t.left = t.left.Add(k, v)
		} else {
			t.right = t.right.Add(k, v)
		}
	}
	// 完成添加, 判断是否需要平衡
	if t.right.Height()-t.left.Height() > 1 {
		if t.right != nil && t.right.left.Height() > t.right.right.Height() {
			// 先对右子树旋转
			t.right.leftRotate()
		}
		t.leftRotate()
		return t
	}
	if t.left.Height()-t.right.Height() > 1 {
		if t.left != nil && t.left.right.Height() > t.left.left.Height() {
			// 如果它的左子树的右子树高度大于它的左子树, 先对左子树左旋
			t.left.leftRotate()
		}
		t.rightRotate()
	}
	return t
}

func (t *Node) InOrder() {
	if t == nil {
		return
	}
	t.left.InOrder()
	fmt.Println(t.k)
	t.right.InOrder()
}

// 左旋
// 1. 创建一个新节点, 值为当前根节点;
// 2. 把新节点的左子树设置为当前节点左子树;
// 3. 把新节点右子树设置为当前节点的右子树的左子树;
// 4. 把当前节点的值设置为右子节点的值;
// 5. 把当前节点的右子树设置为右子树的右子树;
// 6. 把当前节点的左子树设置为新节点;
func (t *Node) leftRotate() {
	newNode := &Node{
		k: t.k,
		v: t.v,
	}
	// step 2
	newNode.left = t.left
	// step 3
	newNode.right = t.right.left
	// step 4
	t.k = t.right.k
	t.v = t.right.v
	// step 5
	t.right = t.right.right
	// step 6
	t.left = newNode
}

// 右旋
// 1. 创建一个新节点, 值等于当前根节点的值;
// 2. 把新节点的右子树设置为当前节点的右子树;
// 3. 把新节点左子树设置为当前节点左子树的右子树;
// 4. 把当前节点的值设置为左子节点的值;
// 5. 把当前节点的左子树设置成左子树的左子树;
// 6. 把当前节点的右子树设置为新节点;
func (t *Node) rightRotate() {
	newNode := &Node{
		k: t.k,
		v: t.v,
	}
	newNode.right = t.right
	newNode.left = t.left.right
	t.k = t.left.k
	t.v = t.left.v
	t.left = t.left.left
	t.right = newNode
}
