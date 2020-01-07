package heap

import "fmt"

type Heap struct {
	data []int
}

func CreateHeap(data []int) *Heap {
	heapData := []int{0}
	h := &Heap{
		data: append(heapData, data...),
	}
	length := len(h.data)
	// 对所有父节点完成堆化
	for i := length/2; i >= 1; i-- {
		h.Heapify(h.data, length, i)
	}
	return h
}

// 堆化, 由每个叶子节点依次向下完成
func (h *Heap) Heapify(arr []int, length int, i int) {
	for {
		maxIdx := i
		if i*2 < length && arr[i] < arr[i*2] {
			maxIdx = i * 2
		}
		if i*2+1 < length && arr[maxIdx] < arr[i*2+1] {
			maxIdx = i*2 + 1
		}
		if maxIdx == i {
			break
		}
		swap(arr, i, maxIdx)
		i = maxIdx
	}
}

// 移除堆顶元素
// 将堆顶元素与最后一个叶子交换, 删除交换后的最后一个叶子
// 将交换后的对顶元素重新调整
func (h *Heap) Pop() int {
	swap(h.data, 1, len(h.data)-1)
	n := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.Heapify(h.data, len(h.data), 1)
	return n
}

func swap(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}

func (h *Heap) Empty() bool {
	return len(h.data) <= 1
}

func (h *Heap) String() string {
	return fmt.Sprintf("data: %v", h.data)
}
