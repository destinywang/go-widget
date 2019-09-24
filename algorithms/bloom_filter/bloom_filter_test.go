package bloom_filter

import (
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
