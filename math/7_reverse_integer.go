package leetcode_math

import "math"

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）
//
// https://leetcode.cn/problems/reverse-integer
func reverse(x int) int {
	var val = 0
	for x != 0 {
		var q = x / 10
		var r = x % 10
		val = val*10 + r
		x = q
	}
	if absInt(val) > math.MaxInt32 {
		return 0
	}
	return val
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
