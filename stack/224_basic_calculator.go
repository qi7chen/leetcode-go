package leetcode_stack

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
// https://leetcode.cn/problems/basic-calculator/

func calculate(s string) int {
	if s == "" {
		return 0
	}
	var sign = 1
	var result int // last result
	var num int    // parsed number
	var stk []int
	for i := 0; i < len(s); i++ {
		var ch = s[i]
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0') // parse number
		} else if ch == '+' || ch == '-' {
			result += sign * num
			if ch == '+' {
				sign = 1
			} else {
				sign = -1
			}
			num = 0
		} else if ch == '(' {
			stk = append(stk, result)
			stk = append(stk, sign)
			result = 0
			sign = 1
		} else if ch == ')' {
			result += sign * num
			num = 0
			result *= stk[len(stk)-1] // mul stack top
			stk = stk[:len(stk)-1]    // pop stack
			result += stk[len(stk)-1] // add stack top
			stk = stk[:len(stk)-1]    // pop stack
		}
	}
	result += sign * num
	return result
}
