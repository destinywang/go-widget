package huffman

import (
	"fmt"
	"sort"
	"strings"
)

func CreateHuffmanTree(content string) *Node {
	contentBytes := []byte(content)
	fmt.Printf("length before compress: %d\n", len(contentBytes))
	
	// init nodes
	nodes := make([]*Node, 0, len(contentBytes))
	existMap := make(map[byte]int)
	for _, b := range contentBytes {
		existMap[b]++
	}
	for k, v := range existMap {
		nodes = append(nodes, &Node{
			data:   k,
			weight: v,
		})
	}
	
	// create huffman tree
	for len(nodes) > 1 {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].weight < nodes[j].weight
		})
		leftNode := nodes[0]
		rightNode := nodes[1]
		nodes = append(nodes[2:], &Node{
			weight: leftNode.weight + rightNode.weight,
			left:   leftNode,
			right:  rightNode,
		})
	}
	
	return nodes[0]
}

// create huffman code
func CreateHuffmanCode(node *Node) map[byte]string {
	return createHuffmanCode0(node, make([]string, 0), make(map[byte]string))
}

func createHuffmanCode0(node *Node, codes []string, codeMp map[byte]string) map[byte]string {
	if node == nil {
		return codeMp
	}
	if node.left == nil && node.right == nil {
		codeMp[node.data] = strings.Join(codes, "")
		return codeMp
	}
	codeMp = createHuffmanCode0(node.left, append(codes, "0"), codeMp)
	codeMp = createHuffmanCode0(node.right, append(codes, "1"), codeMp)
	return codeMp
}

// 压缩
func Compress(content []byte, codeMap map[byte]string) []byte {
	strs := make([]string, 0, len(content))
	for _, b := range content {
		strs = append(strs, codeMap[b])
	}
	huffmanStr := strings.Join(strs, "")
	byteLen := (len(huffmanStr) + 7) / 8
	// create huffman code byte array
	huffmanBytes := make([]byte, 0, byteLen)
	for i := 0; i < len(huffmanStr); i += 8 {
		str := ""
		if i+8 > len(huffmanStr) {
			str = huffmanStr[i:]
		} else {
			str = huffmanStr[i : i+8]
		}
		// str -> byte
		b := Str2Byte(str)
		huffmanBytes = append(huffmanBytes, byte(b))
	}
	return huffmanBytes
}

func Zip(srcBytes []byte) []byte {
	// generic huffman tree
	huffmanTree := CreateHuffmanTree(string(srcBytes))
	// get huffman code map
	codeMap := CreateHuffmanCode(huffmanTree)
	// return
	return Compress(srcBytes, codeMap)
}

// 11100110
func Str2Byte(str string) byte {
	var b byte
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '1' {
			b |= 1 << uint(len(str)-i-1)
		}
	}
	return b
}
