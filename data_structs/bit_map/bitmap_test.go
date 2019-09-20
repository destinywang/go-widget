package bit_map

import (
	"testing"
)

func TestByte2String(t *testing.T) {
	t.Log(byte2String(byte(21)))
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

	t.Logf("10 exist? %v", bm.Exists(10))
	t.Logf("11 exist? %v", bm.Exists(11))
	t.Logf("13 exist? %v", bm.Exists(13))
	bm.Print()
}
