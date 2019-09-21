package huffman

import "strings"

const BitSize = 8

// e.g. 11100110
func Byte2BinString(b byte) string {
	magicNum := 1
	strs := make([]string, 0, BitSize)
	for i := 0; i < BitSize; i++ {
		nm := magicNum << (BitSize - 1 - uint(i))
		if int(b) & magicNum<<(BitSize-1-uint(i)) == nm {
			strs = append(strs, "1")
		} else {
			strs = append(strs, "0")
		}
	}
	return strings.Join(strs, "")
}
// /usr/local/go/src/archive
// /usr/local/go/src/archive