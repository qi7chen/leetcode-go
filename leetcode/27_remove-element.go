package leetcode

// https://leetcode.cn/problems/remove-element

// 原地移除所有数值等于 val 的元素
func removeElement(nums []int, val int) int {
	var slow = 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
