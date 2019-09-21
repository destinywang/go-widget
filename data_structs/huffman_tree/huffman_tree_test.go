package huffman_tree

import "testing"

func TestCreateHuffmanTree(t *testing.T) {
	arr := []int{13, 7, 8, 3, 29, 6, 1}
	t.Log(CreateHuffmanTree(arr).String())
}
