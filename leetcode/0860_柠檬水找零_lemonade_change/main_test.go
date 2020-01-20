package main

import "testing"

func TestLemonadeChange(t *testing.T) {
	t.Log(LemonadeChange([]int{5,5,5,10,20}))
	t.Log(LemonadeChange([]int{5,5,10}))
	t.Log(LemonadeChange([]int{10,20}))
	t.Log(LemonadeChange([]int{5,5,5,5,10,5,10,10,10,20}))
	t.Log(LemonadeChange([]int{5,5,10,10,5,20,5,10,5,5}))
}
