package linked_list

import "testing"

func TestLinkedList_Push(t *testing.T) {
	var linkedList = &LinkedList{}
	linkedList.Append("a")
	linkedList.Append("b")
	linkedList.Append("c")
	linkedList.Append("d")
	linkedList.Append("e")
	t.Log(linkedList.String())
	pos1 := linkedList.Pos("3")
	pos2 := linkedList.Pos("10")
	t.Logf("pos of 3: [%d]", pos1)
	t.Logf("pos of 10: [%d]", pos2)
	linkedList.Del(2)
	t.Log("after del:", linkedList.String())
	linkedList.Insert(2, "new node")
	t.Log("after insert:", linkedList.String())
	linkedList.Reverse()
	t.Log(linkedList.String())
	pop := linkedList.Pop()
	t.Logf("pop=[%s]", pop)
	t.Log(linkedList.String())
}


