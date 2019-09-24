package bloom_filter

import "github.com/DestinyWang/go-widget/data_structs/bit_map"

const defaultSize = 1 << 25

var seeds = []int32{7, 11, 13, 31, 37, 61}

type BloomFilter struct {
	*bit_map.BitMap
	hashs []hash
}

type hash struct {
	cap  int32
	seed int32
}

func (h *hash) hash(value string) int32 {
	var rst int32
	for _, r := range value {
		rst = h.seed + r
	}
	return (h.cap - 1) & rst
}

func (bf *BloomFilter) Init() {
	bf.BitMap = &bit_map.BitMap{
		Size:defaultSize,
	}
	bf.BitMap.Init()
	for i := range seeds {
		bf.hashs = append(bf.hashs, hash{
			cap:defaultSize,
			seed: seeds[i],
		})
	}
}

func (bf *BloomFilter) Add(value string) (err error) {
	for _, h := range bf.hashs {
		if err = bf.BitMap.Set(uint32(h.hash(value))); err != nil {
			return
		}
	}
	return
}

func (bf *BloomFilter) Contains(value string) bool {
	if value == "" {
		return false
	}
	flag := true
	for _, h := range bf.hashs {
		flag = bf.BitMap.Exists(uint32(h.hash(value)))
	}
	return flag
}
