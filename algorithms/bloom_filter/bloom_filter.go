package bloom_filter

import (
	"fmt"
	"github.com/DestinyWang/go-widget/data_structs/bit_map"
	"github.com/sirupsen/logrus"
)

const defaultSize = 1 << 25

var seeds = []int32{7, 11, 13, 31, 37, 61}

type BloomFilter struct {
	bitMap *bit_map.BitMap
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
	bf.bitMap = &bit_map.BitMap{
		Size: defaultSize,
	}
	bf.bitMap.Init()
	for i := range seeds {
		bf.hashs = append(bf.hashs, hash{
			cap:  defaultSize,
			seed: seeds[i],
		})
	}
}

func (bf *BloomFilter) Add(value string) (err error) {
	posList := make([]uint32, 0)
	for _, h := range bf.hashs {
		pos := uint32(h.hash(value))
		posList = append(posList, pos)
		if err = bf.bitMap.Set(pos); err != nil {
			logrus.WithError(err).Errorf("set bit map fail")
			return
		}
	}
	fmt.Printf("value=[%s] posList=[%v]\n", value, posList)
	return
}

func (bf *BloomFilter) Contains(value string) bool {
	if value == "" {
		return false
	}
	flag := true
	for _, h := range bf.hashs {
		flag = bf.bitMap.Exists(uint32(h.hash(value)))
	}
	return flag
}
