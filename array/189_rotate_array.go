package leetcode_array

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
// https://leetcode.cn/problems/rotate-array/

// 当我们将数组的元素向右移动 k 次后，尾部 k % n 个元素会移动至数组头部，其余元素向后移动 k % n 个位置。
// 该方法为数组的翻转：我们可以先将所有元素翻转，这样尾部的 k%n 个元素就被移至数组头部，
// 然后我们再翻转 [0,k%n−1] 区间的元素和 [k%n,n−1] 区间的元素即能得到最后的答案

func rotate(nums []int, k int) {
	if len(nums) <= 1 || k <= 0 {
		return
	}
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 利用额外空间
func rotateV1(nums []int, k int) {
	if len(nums) <= 1 || k <= 0 {
		return
	}
	k %= len(nums)
	var buf = make([]int, k)
	copy(buf, nums[len(nums)-k:])
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = buf[i]
	}
}
