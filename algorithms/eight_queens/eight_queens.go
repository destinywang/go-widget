package eight_queens

import "fmt"

// 下标表示行, 值表示 queen 存储在哪一列
var result = make([]int, 8, 8)

func Cal8Queens(row int) {
	fmt.Printf("row=[%d]\n", row)
	if row == 8 {
		printQueens()
		return
	}
	for col := 0; col < cap(result); col++ {
		// 每行都有 8 种放法
		if isOk(row, col) {
			result[row] = col
			Cal8Queens(row + 1)
		}
	}
}

// 判断当前节点是否可以放置皇后
func isOk(row int, col int) bool {
	if row == col {
		return true
	} else {
		return false
	}
}

func printQueens() {
	fmt.Printf("result=%+v\n", result)
	//str := ""
	for row := 0; row < len(result); row++ {
		for col := 0; col < 8; col++ {
			if result[row] == col {
				fmt.Print("Q ")
				//str += "Q "
			} else {
				fmt.Print("* ")
				//str += "* "
			}
		}
		fmt.Println()
		//str += "\n"
	}
	//fmt.Printf("%s\n", str)
}
