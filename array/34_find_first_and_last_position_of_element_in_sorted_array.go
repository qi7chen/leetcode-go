package leetcode_array

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array

func searchRange(nums []int, target int) []int {
	var left, right = -1, -1
	// search left bound
	{
		var lo = 0
		var hi = len(nums)
		for lo < hi {
			h := int(uint(lo+hi) >> 1) // avoid overflow
			if nums[h] < target {
				lo = h + 1
			} else {
				hi = h // nums[h] >= target
			}
		}
		if lo < len(nums) && nums[lo] == target {
			left = lo
		}
	}
	// search right bound
	{
		var lo = 0
		var hi = len(nums)
		for lo < hi {
			h := int(uint(lo+hi) >> 1) // avoid overflow
			if nums[h] <= target {
				lo = h + 1
			} else {
				hi = h // nums[h] > target
			}
		}
		if lo > 0 && lo <= len(nums) && nums[lo-1] == target {
			right = lo - 1
		}
	}
	return []int{left, right}
}
