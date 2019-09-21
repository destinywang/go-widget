package huffman_code

import (
	"testing"
)

func TestCreateHuffmanCode(t *testing.T) {
	s := "Over-time the usage of the log grew from an implementation detail of ACID to a method for replicating data between databases. It turns out that the sequence of changes that happened on the database is exactly what is needed to keep a remote replica database in sync. Oracle, MySQL, and PostgreSQL include log shipping protocols to transmit portions of log to replica databases which act as slaves. Oracle has productized the log as a general data subscription mechanism for non-oracle data subscribers with their XStreams and GoldenGate and similar facilities in MySQL and PostgreSQL are key components of many data architectures."
	node := CreateHuffmanTree(s)
	t.Log(node)
	codeMap := CreateHuffmanCode(node)
	t.Log(codeMap)
	compress := Compress([]byte(s), codeMap)
	t.Log(compress)
}

func TestParseInt(t *testing.T) {
	//i, err := strconv.ParseInt("11100110", 2, 8)
	//assert.Nil(t, err)
	//t.Log(i)
	t.Log(1 <<0)
}

func TestStr2Byte(t *testing.T) {
	b := Str2Byte("11100110")
	t.Log(b)
}
