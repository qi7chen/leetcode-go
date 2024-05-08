package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	var head = &ListNode{Val: s[0]}
	var prev = head
	for i := 1; i < len(s); i++ {
		var node = &ListNode{Val: s[i]}
		prev.Next = node
		prev = node
	}
	return head
}

func fmtList(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var s []int
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	return s
}

func isEqualList(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func countList(root *ListNode) int {
	var count int
	for root != nil {
		count++
		root = root.Next
	}
	return count
}

func toInt(v any) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	}
	return 0, fmt.Errorf("invalid type %T", v)
}

func toIntList(s []any) []int {
	var a []int
	for _, v := range s {
		val, err := toInt(v)
		if err == nil {
			a = append(a, val)
		}
	}
	return a
}
