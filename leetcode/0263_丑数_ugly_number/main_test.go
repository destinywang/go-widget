package main

import "testing"

func TestIsUgly(t *testing.T) {
	t.Log(IsUgly(6))
	t.Log(IsUgly(8))
	t.Log(IsUgly(14))
}
