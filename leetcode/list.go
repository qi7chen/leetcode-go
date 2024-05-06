package leetcode

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

func equalList(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
