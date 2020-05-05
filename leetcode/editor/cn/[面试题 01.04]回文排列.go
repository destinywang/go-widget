package cn

//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
//
// 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。 
//
// 回文串不一定是字典当中的单词。 
//
// 
//
// 示例1： 
//
// 输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）
// 
//
// 
// Related Topics 哈希表 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	var charNumMap = make(map[int32]int)
	for _, r := range s {
		charNumMap[r]++
	}
	var flag bool
	if len(s)%2 == 0 {
		// 如果字符数是偶数, 所有字符的个数都必须是偶数
		flag = true
		for _, v := range charNumMap {
			if v%2 != 0 {
				flag = false
			}
		}
	} else {
		flag = false
		var p1 int
		for _, v := range charNumMap {
			if v%2 != 0 {
				p1++
			}
		}
		if p1 == 1 {
			flag = true
		}
	}
	return flag
}

//leetcode submit region end(Prohibit modification and deletion)
