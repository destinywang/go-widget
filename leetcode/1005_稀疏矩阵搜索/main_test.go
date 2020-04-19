package _005_稀疏矩阵搜索

import "testing"

func TestFindString(t *testing.T) {
	words := []string{
		"at",
		"",
		"",
		"",
		"ball",
		"",
		"",
		"cat",
		"",
		"",
		"dad",
		"",
		"",
	}
	t.Log(FindStringByBinSearch(words, "ta"))
	t.Log(FindStringByBinSearch(words, "ball"))
}
