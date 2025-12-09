package leetcode_array

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
// nums1 的初始长度为 m + n，后 n 个元素为 0
// https://leetcode.cn/problems/merge-sorted-array

// 将两个数组看作队列，每次从两个数组尾部取出比较大的数字放到结果中 O(m+n)
func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	var i = m - 1
	var j = n - 1
	var k = m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
