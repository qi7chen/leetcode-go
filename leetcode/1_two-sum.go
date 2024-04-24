package leetcode

// https://leetcode.cn/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	return twosumV2(nums, target)
}

func twosumV1(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
	}
	return []int{-1, -1}
}

// hash表 O(N)
func twosumV2(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
	}
	return []int{-1, -1}
}

// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// 根据数值范围，用数值替换map
func twosumV3(nums []int, target int) []int {
	const MaxN = 109
	var table [MaxN*2 + 1]int
	for i, num := range nums {
		table[num+MaxN] = i + 1
	}
	for i, num := range nums {
		var j = table[target-num+MaxN]
		if j > 0 {
			return []int{j - 1, i}
		}
	}
	return []int{-1, -1}
}
