package leetcode_array

// 用2个指针，一个在前记作 p，一个在后记作 q，算法流程如下：
// 比较 p 和 q 位置的元素是否相等。
// 如果相等，q 后移 1 位
// 如果不相等，将 q 位置的元素复制到 p+1 位置上，p 后移一位，q 后移 1 位
// 重复上述过程，直到 q 等于数组长度。
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var p, q = 0, 1
	for q < len(nums) {
		if nums[p] != nums[q] {
			p++
			nums[p] = nums[q]
		}
		q++
	}
	return p + 1
}
