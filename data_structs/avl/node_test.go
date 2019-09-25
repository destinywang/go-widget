package avl

import "testing"

func TestTreeNode_Height(t *testing.T) {
	tn := &Node{
		k: 1,
		left: &Node{
			k: 2,
			left: &Node{
				k: 3,
				left: &Node{
					k: 4,
				},
			},
		},
		right: &Node{
			k: 5,
			right: &Node{
				k: 6,
				left: &Node{
					k: 7,
					left: &Node{
						k: 8,
					},
				},
			},
		},
	}
	t.Log(tn.Height())
	tn.InOrder()
}

func TestTreeNode_Add(t *testing.T) {
	at := new(Tree)
	arr := []int{10, 11, 7, 6, 8, 9}
	for _, i := range arr {
		at.Add(i, i)
	}
	t.Log(at.Height())
	at.InOrder()
	t.Logf("left sub tree: [%d]", at.LeftHeight())
	t.Logf("right sub tree: [%d]", at.RightHeight())
}
