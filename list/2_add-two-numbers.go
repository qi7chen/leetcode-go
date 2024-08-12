package leetcode_list

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// https://leetcode.cn/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head = l1
	var prev *ListNode
	for l1 != nil || l2 != nil {
		var node = l1
		if node == nil {
			node = l2
		} else if l2 != nil {
			node.Val += l2.Val
		}
		node.Val += carry
		carry = 0
		if node.Val >= 10 {
			carry = 1
			node.Val -= 10 // 0 <= Val <= 9
		}
		if prev != nil {
			prev.Next = node
		}
		prev = node // 合并节点，不做分配
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		prev.Next = &ListNode{Val: carry}
	}
	return head
}
