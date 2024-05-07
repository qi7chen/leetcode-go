package leetcode

import (
	"slices"
)

// https://leetcode.cn/problems/merge-sorted-array

// O(m+n)
// 将两个数组看作队列，每次从两个数组尾部取出比较大的数字放到结果中
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

// O((m+n)log(m+n))
// 将数组arr放入dst的尾部，然后直接对整个数组进行排序
func mergeV2(dst []int, m int, arr []int, n int) {
	for i := 0; i < n; i++ {
		dst[m+i] = arr[i]
	}
	slices.Sort(dst)
}
