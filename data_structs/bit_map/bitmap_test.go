package bit_map

import (
	"testing"
)

func TestByte2String(t *testing.T) {
	t.Log(Byte2String(byte(21)))
}

func TestBitMap_Set(t *testing.T) {
	bm := &BitMap{
		size: 20,
	}
	bm.Init()
	bm.Print()
	bm.Set(10)
	bm.Set(13)
	bm.Set(15)
	bm.Set(17)
	bm.Set(19)
	bm.Print()
}
