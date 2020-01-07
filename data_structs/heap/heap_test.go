package heap

import (
	"testing"
)

func TestCreateHeap(t *testing.T) {
	//var data []int
	//for i := 0; i < 10; i++ {
	//	data = append(data, randomdata.Number(1, 100))
	//}
	data := []int{18,66,50,49,96,44,25,72,17,54}
	t.Log("origin: ", data)
	heap := CreateHeap(data)
	t.Log(heap)
	t.Log(heap.Pop())
	t.Log(heap)
}
