package leetcode_string

import (
	"math"
)

// https://leetcode.cn/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	// skip leading space
	var i = 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	var neg bool
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		neg = s[i] == '-'
		i++
	}
	var val int64
	for i < len(s) && isDigit(s[i]) {
		val = val*10 + (int64(s[i]) - int64('0'))
		if val > math.MaxInt32 {
			if neg {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		i++
	}
	if neg {
		val = max(-val, math.MinInt32)
	} else {
		val = min(val, math.MaxInt32)
	}
	return int(val)
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
