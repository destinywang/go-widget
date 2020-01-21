package main

import (
	"fmt"
	"github.com/DestinyWang/go-widget/leetcode"
	"strconv"
)

// 你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。
// 空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

func Tree2Str(t *leetcode.TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return strconv.FormatInt(int64(t.Val), 10)
	}
	if t.Right == nil {
		return fmt.Sprintf("%s(%s)", strconv.FormatInt(int64(t.Val), 10), Tree2Str(t.Left))
	}
	return fmt.Sprintf("%s(%s)(%s)", strconv.FormatInt(int64(t.Val), 10), Tree2Str(t.Left), Tree2Str(t.Right))
}

func main() {

}
