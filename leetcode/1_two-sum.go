package leetcode

// https://leetcode.cn/problems/two-sum/description/
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 使用hash表 O(N)
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		if j, ok := m[target-num]; ok && j != i {
			return []int{j, i}
		}
	}
	return []int{-1, -1}
}
