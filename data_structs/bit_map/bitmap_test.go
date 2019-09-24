package bit_map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByte2String(t *testing.T) {
	t.Log(byte2String(byte(21)))
}

func TestBitMap_Set(t *testing.T) {
	bm := &BitMap{
		Size: 20,
	}
	bm.Init()
	t.Logf("start: [%s]", bm)
	err := bm.Set(10)
	assert.Nil(t, err)
	err = bm.Set(13)
	assert.Nil(t, err)
	err = bm.Set(99)
	assert.NotNil(t, err)

	exists := bm.Exists(10)
	assert.True(t, exists)
	exists = bm.Exists(11)
	assert.False(t, exists)
	exists = bm.Exists(13)
	assert.True(t, exists)
	t.Logf("end: [%s]", bm)
}
