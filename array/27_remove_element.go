package leetcode_array

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。
// 然后返回 nums 中与 val 不同的元素的数量
// https://leetcode.cn/problems/remove-element

// 左右双指针解法
func removeElement(nums []int, val int) int {
	var left, right = 0, 0
	for right < len(nums) {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}

// 双指针优化
func removeElementDoublePointer(nums []int, val int) int {
	var left = 0
	var right = len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
