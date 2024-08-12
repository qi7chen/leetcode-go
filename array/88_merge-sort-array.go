package leetcode_array

import (
	"slices"
)

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
// https://leetcode.cn/problems/merge-sorted-array

// 将两个数组看作队列，每次从两个数组尾部取出比较大的数字放到结果中 O(m+n)
func merge(dst []int, m int, arr []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if dst[i] > arr[j] {
			dst[k] = dst[i]
			k--
			i--
		} else {
			dst[k] = arr[j]
			k--
			j--
		}
	}
	for j >= 0 {
		dst[k] = arr[j]
		j--
		k--
	}
}

// 将数组arr放入dst的尾部，然后直接对整个数组进行排序
// O((m+n)log(m+n))
func mergeV2(dst []int, m int, arr []int, n int) {
	for i := 0; i < n; i++ {
		dst[m+i] = arr[i]
	}
	slices.Sort(dst)
}
