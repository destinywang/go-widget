package _005_稀疏矩阵搜索

// hash 表思路
func FindString(words []string, s string) int {
	wordMap := make(map[string]int)
	for i := range words {
		if words[i] != "" {
			wordMap[words[i]] = i
		}
	}
	if i, ok := wordMap[s]; ok {
		return i
	} else {
		return -1
	}
}
