package leetcode_array

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// https://leetcode.cn/problems/search-insert-position

// 二分查找 https://mp.weixin.qq.com/s/M1KfTfNlu4OCK8i9PSAmug
func searchInsert(nums []int, target int) int {
	var left = 0
	var right = len(nums)
	for left < right {
		h := int(uint(left+right) >> 1) // avoid overflow
		if nums[h] < target {
			left = h + 1 // left bound
		} else {
			right = h // nums[m] > target
		}
	}
	return left
}
