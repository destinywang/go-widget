package bloom_filter

import (
	hash2 "github.com/DestinyWang/go-widget/algorithms/hash"
	"github.com/DestinyWang/go-widget/data_structs/bit_map"
	"github.com/sirupsen/logrus"
)

const defaultSize = 1<<32 - 1

var seeds = []int64{7, 11, 13, 31, 37, 61}

type BloomFilter struct {
	bitMap *bit_map.BitMap
	hashs  []hash
}

type hash struct {
	seed int64
}

func (h *hash) hash(value string) int64 {
	return hash2.MurmurHash64A([]byte(value), h.seed)
}

func (bf *BloomFilter) Init() {
	bf.bitMap = &bit_map.BitMap{
		Size: defaultSize,
	}
	bf.bitMap.Init()
	for i := range seeds {
		bf.hashs = append(bf.hashs, hash{
			seed: seeds[i],
		})
	}
}

func (bf *BloomFilter) Add(value string) (err error) {
	//posList := make([]uint32, 0)
	for _, h := range bf.hashs {
		pos := uint32(h.hash(value)) % (1 << 31)
		//posList = append(posList, pos)
		if err = bf.bitMap.Set(pos); err != nil {
			logrus.WithError(err).Errorf("set bit map fail")
			return
		}
	}
	//fmt.Printf("value=[%s] posList=[%v]\n", value, posList)
	return
}

func (bf *BloomFilter) Contains(value string) bool {
	if value == "" {
		return false
	}
	flag := true
	for _, h := range bf.hashs {
		flag = flag && bf.bitMap.Exists(uint32(h.hash(value))%(1<<31))
	}
	return flag
}
