package bloom_filter

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := &BloomFilter{}
	bf.Init()
	contains := bf.Contains("destiny")
	t.Logf("before: %t", contains)
	err := bf.Add("destiny")
	assert.Nil(t, err)
	contains = bf.Contains("destiny")
	t.Logf("after: %t", contains)
}

func Test(t *testing.T) {
	bf := &BloomFilter{}
	bf.Init()
	cache := make(map[string]struct{})
	errCnt := 0
	loopCnt := 50000000
	for i := 0; i < loopCnt; i++ {
		//uuids := uuid.NewV4()
		uid := strconv.FormatInt(int64(i), 10)
		inBloom := bf.Contains(uid)
		_, inMap := cache[uid]
		if inBloom != inMap {
			t.Logf("idx=[%d], uuid=[%s], inBloom=[%v], inMap=[%v]", i, uid, inBloom, inMap)
			errCnt++
		}
		cache[uid] = struct{}{}
		err := bf.Add(uid)
		assert.Nil(t, err)
	}
	assert.Equal(t, loopCnt, len(cache))
	assert.True(t, bf.Contains("100"))
	assert.True(t, bf.Contains("1000"))
	assert.True(t, bf.Contains("10000"))
	assert.True(t, bf.Contains("100000"))
	assert.True(t, bf.Contains("1000000"))
	assert.True(t, bf.Contains("10000000"))
	assert.False(t, bf.Contains(strconv.FormatInt(int64(loopCnt), 10)))
	t.Logf("errCnt=[%d]", errCnt)
}
