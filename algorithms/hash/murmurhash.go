package hash

const (
	BigM = 0xc6a4a7935bd1e995
	BigR = 47
)

func MurmurHash64A(data []byte, seed int64) (h int64) {
	var k int64
	h = seed ^ int64(uint64(len(data))*BigM)

	var ubigm uint64 = BigM
	var ibigm = int64(ubigm)
	for l := len(data); l >= 8; l -= 8 {
		k = int64(data[0]) | int64(data[1])<<8 | int64(data[2])<<16 | int64(data[3])<<24 | int64(data[4])<<32 | int64(data[5])<<40 | int64(data[6])<<48 | int64(data[7])<<56

		k := k * ibigm
		k ^= int64(uint64(k) >> BigR)
		k = k * ibigm

		h = h ^ k
		h = h * ibigm
		data = data[8:]
	}

	switch len(data) {
	case 7:
		h ^= int64(data[6]) << 48
		fallthrough
	case 6:
		h ^= int64(data[5]) << 40
		fallthrough
	case 5:
		h ^= int64(data[4]) << 32
		fallthrough
	case 4:
		h ^= int64(data[3]) << 24
		fallthrough
	case 3:
		h ^= int64(data[2]) << 16
		fallthrough
	case 2:
		h ^= int64(data[1]) << 8
		fallthrough
	case 1:
		h ^= int64(data[0])
		h *= ibigm
	}

	h ^= int64(uint64(h) >> BigR)
	h *= ibigm
	h ^= int64(uint64(h) >> BigR)
	return
}
