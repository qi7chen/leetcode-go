package leetcode_array

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// https://leetcode.cn/problems/two-sum/
// 使用hash表 O(N)
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[target-num]; ok && j != i {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

// 暴力解法 O(N^2)
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
