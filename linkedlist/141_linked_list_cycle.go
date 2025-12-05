package leetcode_list

// 判断链表中是否有环
// https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true // 慢指针和快指针相遇，说明有环
		}
	}
	return false
}
