package leetcode_array

// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
// 返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array

// 双指针解法
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var slow, fast = 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
