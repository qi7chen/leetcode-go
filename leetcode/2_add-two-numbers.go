package leetcode

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
