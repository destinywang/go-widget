package bit_map

import (
	"fmt"
	"math"
	"strings"
)

const BitSize = 8

type BitMap struct {
	bitArray []byte
	Size     uint32
}

func (bm *BitMap) Init() {
	if bm.Size > 1<<27 {
		panic("bit map Size over flow")
	}
	var r uint32
	if bm.Size <= 1 {
		r = 1
	} else {
		r = uint32(math.Ceil(float64(float64(bm.Size) / BitSize)))
	}
	bm.bitArray = make([]byte, r)
}

func (bm *BitMap) calIdx(i uint32) (idx uint32, pos uint32) {
	idx = i >> 3
	pos = i & (BitSize - 1)
	return
}

func (bm *BitMap) Set(i uint32) error {
	if i > bm.Size {
		return fmt.Errorf("bit map Size over flow, max Size=[%d]", bm.Size)
	}
	idx, pos := bm.calIdx(i)
	bm.bitArray[idx] |= 1 << (pos - 1)
	return nil
}

func (bm *BitMap) Exists(i uint32) bool {
	if i > bm.Size {
		return false
	}
	idx, pos := bm.calIdx(i)
	return bm.bitArray[idx]&(1<<(pos-1)) == 1<<(pos-1)
}

func (bm *BitMap) String() string {
	strs := make([]string, 0, len(bm.bitArray))
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
