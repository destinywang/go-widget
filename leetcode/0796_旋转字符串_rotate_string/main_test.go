package main

import "testing"

func TestRotateString(t *testing.T) {
	t.Log(RotateString("abcde", "cdeab"))
	t.Log(RotateString("abcde", "abced"))
}
