package bloom_filter

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
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

	for i := 0; i < 10000000; i++ {
		uuids := uuid.NewV4()
		uid := uuids.String()
		inBloom := bf.Contains(uid)
		_, inMap := cache[uid]
		if inBloom != inMap {
			t.Errorf("idx=[%d], uuid=[%s], inBloom=[%v], inMap=[%v]", i, uid, inBloom, inMap)
		}

		cache[uid] = struct{}{}
		bf.Add(uid)
	}
}
