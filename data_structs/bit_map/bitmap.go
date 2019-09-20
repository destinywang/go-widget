package bit_map

import (
	"fmt"
	"math"
	"strings"
)

const BitSize = 8

type BitMap struct {
	bitArray []byte
	size     uint32
}

func (bm *BitMap) Init() {
	var r uint32
	if bm.size <= 1 {
		r = 1
	} else {
		r = uint32(math.Ceil(float64(float64(bm.size) / BitSize)))
	}
	bm.bitArray = make([]byte, r)
}

func (bm *BitMap) calIdx(i uint32) (idx uint32, pos uint32) {
	idx = i >> 3
	pos = i & (BitSize - 1)
	return
}

func (bm *BitMap) Set(i uint32) error {
	if i > bm.size {
		return fmt.Errorf("bit map size over flow, max size=[%d]", bm.size)
	}
	idx, pos := bm.calIdx(i)
	bm.bitArray[idx] |= 1 << (pos-1)
	return nil
}

func (bm *BitMap) Exists(i uint32) bool {
	if i > bm.size {
		return false
	}
	idx, pos := bm.calIdx(i)
	return bm.bitArray[idx]&(1<<(pos-1)) == 1<<(pos-1)
}

func (bm *BitMap) String() string {
	strs := make([]string, len(bm.bitArray))
	for _, b := range bm.bitArray {
		strs = append(strs, byte2String(b))
	}
	return strings.Join(strs, "-")
}

func byte2String(b byte) string {
	var magicBit uint32 = 1
	str := ""
	for i := 0; i < BitSize; i++ {
		u := magicBit << uint32(i)
		if uint32(b)&u == u {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}
