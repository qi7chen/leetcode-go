package leetcode_list

// 给你两个有序链表，合并两个链表并使新链表中的节点仍然有序。
// https://leetcode.cn/problems/merge-two-sorted-lists/description/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy ListNode
	var node = &dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}
	if list1 == nil {
		node.Next = list2
	} else {
		node.Next = list1
	}
	return dummy.Next
}
