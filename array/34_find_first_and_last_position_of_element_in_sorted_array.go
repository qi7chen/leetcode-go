package leetcode_array

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array

func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, func(i int) bool {
		return nums[i] < target
	})
	if !(left < len(nums) && nums[left] == target) {
		left = -1
	}
	right := binarySearch(nums, func(i int) bool {
		return nums[i] <= target
	})
	right -= 1
	if !(right >= 0 && right < len(nums) && nums[right] == target) {
		right = -1
	}
	return []int{left, right}
}

func binarySearch(x []int, cmp func(int) bool) int {
	var left = 0
	var right = len(x)
	for left < right {
		h := int(uint(left+right) >> 1) // avoid overflow
		if cmp(h) {
			left = h + 1
		} else {
			right = h //  x[h] >= target
		}
	}
	return left
}

// leftBound = binarySearch(nums[i] < target)
func leftBound(x []int, target int) int {
	var left = 0
	var right = len(x)
	for left < right {
		h := int(uint(left+right) >> 1) // avoid overflow
		if x[h] < target {
			left = h + 1
		} else {
			right = h // x[h] >= target
		}
	}
	return left
}

// rightBound = binarySearch(nums[i] <= target) - 1
func rightBound(x []int, target int) int {
	var left = 0
	var right = len(x)
	for left < right {
		h := int(uint(left+right) >> 1) // avoid overflow
		if x[h] <= target {
			left = h + 1
		} else {
			right = h // x[h] > target
		}
	}
	return left - 1
}
