package leetcode_list

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表
// https://leetcode.cn/problems/reverse-linked-list/
func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	for node != nil {
		var next = node.Next
		node.Next = prev
		prev, node = node, next
	}
	return prev
}
