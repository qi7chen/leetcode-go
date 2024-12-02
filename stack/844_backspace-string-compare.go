package leetcode_stack

import (
	"unsafe"
)

// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符
// https://leetcode.cn/problems/backspace-string-compare/

func processBackspaceInput(stk []byte, s string) []byte {
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1] // pop stack
			}
		} else {
			stk = append(stk, s[i]) // push stack
		}
	}
	return stk
}

func backspaceCompare(s, t string) bool {
	var stk1 = make([]byte, 0, len(s))
	var stk2 = make([]byte, 0, len(t))

	stk1 = processBackspaceInput(stk1, s)
	stk2 = processBackspaceInput(stk2, t)

	s1 := unsafe.String(unsafe.SliceData(stk1), len(stk1))
	s2 := unsafe.String(unsafe.SliceData(stk2), len(stk2))
	return s1 == s2
}
