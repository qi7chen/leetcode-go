package leetcode

import "math"

// https://leetcode.cn/problems/reverse-integer
func reverse(x int) int {
	var val = 0
	for x != 0 {
		var q = x / 10
		var r = x % 10
		val = val*10 + r
		x = q
	}
	if abs(val) > math.MaxInt32 {
		return 0
	}
	return val
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
