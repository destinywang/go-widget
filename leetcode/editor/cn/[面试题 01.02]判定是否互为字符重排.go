package cn
//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
//
// 示例 1： 
//
// 输入: s1 = "abc", s2 = "bca"
//输出: true 
// 
//
// 示例 2： 
//
// 输入: s1 = "abc", s2 = "bad"
//输出: false
// 
//
// 说明： 
//
// 
// 0 <= len(s1) <= 100 
// 0 <= len(s2) <= 100 
// 
// Related Topics 数组 字符串


//leetcode submit region begin(Prohibit modification and deletion)
func CheckPermutation(s1 string, s2 string) bool {
	var charMap = make(map[int32]int)
	for _, r := range s1 {
		charMap[r]++
	}
	for _, r := range s2 {
		if charMap[r] == 0 {
			return false
		}
		charMap[r]--
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
