package leetcode_array

import (
	"sort"
)

// 给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
// 换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
// 以数组形式返回答案。
// https://leetcode.cn/problems/how-many-numbers-are-smaller-than-the-current-number/

func smallerNumbersThanCurrent(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var sorted = make([]Sorter, len(nums))
	for i := 0; i < len(nums); i++ {
		sorted[i] = Sorter{val: nums[i], idx: i}
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].val < sorted[j].val
	})

	var prev = -1
	var ans = make([]int, len(nums))
	for i := 0; i < len(sorted); i++ {
		if prev == -1 || sorted[i].val != sorted[i-1].val {
			prev = i
		}
		ans[sorted[i].idx] = prev
	}
	return ans
}

type Sorter struct {
	val int
	idx int
}
