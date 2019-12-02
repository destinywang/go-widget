package lru

import "testing"

func TestDefaultLru_Push(t *testing.T) {
	lru := NewDefaultLru(5)
	lru.Push("a")
	lru.Push("b")
	lru.Push("c")
	lru.Push("d")
	lru.Push("e")
	lru.Push("f")
	t.Log(lru.String())
}
