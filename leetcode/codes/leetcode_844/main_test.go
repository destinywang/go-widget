package main

import (
	"testing"
)

func TestInitStack(t *testing.T) {
	t.Log(InitStack("y#fo##f"))
	t.Log(InitStack("y#f#o##f"))
}