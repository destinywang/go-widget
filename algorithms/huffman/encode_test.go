package huffman

import (
	//"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCreateHuffmanCode(t *testing.T) {
	//s := "Over-time the usage of the log grew from an implementation detail of ACID to a method for replicating data between databases. It turns out that the sequence of changes that happened on the database is exactly what is needed to keep a remote replica database in sync. Oracle, MySQL, and PostgreSQL include log shipping protocols to transmit portions of log to replica databases which act as slaves. Oracle has productized the log as a general data subscription mechanism for non-oracle data subscribers with their XStreams and GoldenGate and similar facilities in MySQL and PostgreSQL are key components of many data architectures."
	//s := "i like like like java do you like a java"
	// ilo ilo iv kkkkeyio ilo iv kkkkeyio ilo iv kkkkeyiv eja eda iv eda a eda ilo iv eyio io a uillo ilo iv kkkkeyiv da a eja eda illlv e
	s := "aaabbc"
	node := CreateHuffmanTree(s)
	//t.Log(node)
	codeMap := CreateHuffmanCode(node)
	t.Log(codeMap)
	compress := Encode([]byte(s), codeMap)
	//t.Logf("compress=[%v]", compress)
	//fmt.Printf("compress=[%v]\n", compress)
	Decode(compress, codeMap)
}

func TestParseInt(t *testing.T) {
	i, _ := strconv.ParseInt("11100110", 2, 16)
	//assert.Nil(t, err)
	t.Log(i)
	//atoi, _ := strconv.Atoi("11100110")
	//t.Log(atoi)
	//t.Log(1 <<0)
}

func TestStr2Byte(t *testing.T) {
	b := Str2Byte("11100110")
	t.Log(b)
}

func TestHuffmanZip(t *testing.T) {
	s := "i like like like java do you like a java"
	bytes := Zip([]byte(s))
	t.Logf("length of huffman code=%d", len(bytes))
}
