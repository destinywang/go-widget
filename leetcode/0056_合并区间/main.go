package _056_合并区间

import (
	"math"
	"sort"
)

// 给出一个区间的集合, 请合并所有重叠的区间

func Merge(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var mergeFlag = make(map[int]struct{})
	var merged [][]int
	for i := 0; i < len(intervals); i++ {
		if _, ok := mergeFlag[i]; ok {
			continue
		}
		s := []int{intervals[i][0], intervals[i][1]}
		for j := i; j < len(intervals); j++ {
			if s[1] >= intervals[j][0] {
				s[0] = min(s[0], intervals[j][0])
				s[1] = max(s[1], intervals[j][1])
				mergeFlag[j] = struct{}{}
			}
		}
		merged = append(merged, s)
	}
	return merged
}

func max(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}

func min(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}

// 先通过排序降低复杂度, 排序后两个集合只要 intervals[i][1] <= intervals[j][0] 时就满足包含的条件, 上下边界各取最小值和最大值即可
