package huffman

import (
	"fmt"
	"strings"
	"time"
)

const BitSize = 8

// e.g. 11100110
func Byte2BinString(b byte) string {
	magicNum := 1
	strs := make([]string, 0, BitSize)
	for i := 0; i < BitSize; i++ {
		nm := magicNum << (BitSize - 1 - uint(i))
		if int(b)&nm == nm {
			strs = append(strs, "1")
		} else {
			strs = append(strs, "0")
		}
	}
	return strings.Join(strs, "")
}

// 1. Huffman code bytes => Huffman binary string
// 2. Huffman binary string => origin string
func Decode(huffmanBytes []byte, codeMap map[byte]string) []byte {
	binStrs := make([]string, 0, len(huffmanBytes))
	for _, b := range huffmanBytes {
		binStrs = append(binStrs, Byte2BinString(b))
	}
	huffmanBinStr := strings.Join(binStrs, "")
	fmt.Printf("[Decode] huffmanBytes=[%v]\n", huffmanBytes)
	fmt.Printf("[Decode] huffmanBinStr=[%s]\n", huffmanBinStr)
	// revert huffmanCodeMap
	huffmanCodeMap := make(map[string]byte)
	for k, v := range codeMap {
		huffmanCodeMap[v] = k
	}
	src := make([]byte, 0)
	for i := 0; i < len(huffmanBinStr); {
		for j := i + 1; j < len(huffmanBinStr); {
			if b, ok := huffmanCodeMap[huffmanBinStr[i:j]]; ok {
				//fmt.Printf("b=[%s]", string(b))
				src = append(src, b)
				i = j
				break
			} else {
				j++
			}
			if j == len(huffmanBinStr) {
				i = j
			}
		}
		time.Sleep(10 * time.Millisecond)
		//fmt.Printf("i=[%d]\n", i)
	}
	//fmt.Printf("[Decode] src=[%v]", src)
	fmt.Printf("[Decode] src=[%s]\n", string(src))
	return nil
}
