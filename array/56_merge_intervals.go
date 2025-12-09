package leetcode_array

import (
	"sort"
)

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
// https://leetcode.cn/problems/merge-intervals/description/

func mergeInterval(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] // 按左端点排序
	})
	var result = make([][]int, 0, len(intervals))
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		var last = result[len(result)-1]
		var current = intervals[i]
		if current[0] <= last[1] { // 有重叠
			if current[1] > last[1] {
				last[1] = current[1] // 合并区间
			}
		} else {
			result = append(result, current)
		}
	}
	return result
}
