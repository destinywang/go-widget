package _997_找到小镇的法官_find_the_town_judge

import "testing"

func TestFindJudge(t *testing.T) {
	judgeIdx := FindJudge(4, [][]int{
		{
			1, 3,
		},
		{
			1, 4,
		},
		{
			2, 3,
		},
		{
			2, 4,
		},
		{
			4, 3,
		},
	})
	t.Log(judgeIdx)
}
