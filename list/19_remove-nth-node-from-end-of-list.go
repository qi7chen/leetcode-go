package leetcode_list

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{Next: head}
	var second = dummy // 右指针先移动N步
	for i := 0; i <= n; i++ {
		second = second.Next
	}
	var first = dummy // 左右指针同时移动
	for second != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next // 删除
	return dummy.Next
}
