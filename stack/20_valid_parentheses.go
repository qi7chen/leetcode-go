package leetcode_stack

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// https://leetcode.cn/problems/valid-parentheses/

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])

		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1] // pop stack

		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1] // pop stack

		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1] // pop stack
		}
	}
	return len(stack) == 0
}
