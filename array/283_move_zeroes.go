package leetcode_array

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 必须在不复制数组的情况下原地对数组进行操作
// https://leetcode.cn/problems/move-zeroes

func moveZeroes(nums []int) {
	var pos int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}
	clear(nums[pos:])
}
