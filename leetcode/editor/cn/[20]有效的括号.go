package cn

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 注意空字符串可被认为是有效字符串。 
//
// 示例 1: 
//
// 输入: "()"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "()[]{}"
//输出: true
// 
//
// 示例 3: 
//
// 输入: "(]"
//输出: false
// 
//
// 示例 4: 
//
// 输入: "([)]"
//输出: false
// 
//
// 示例 5: 
//
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	var pattMap = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	for _, r := range s {
		if _, ok := pattMap[r]; !ok {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			if pattMap[r] != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
