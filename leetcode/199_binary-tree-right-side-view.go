package leetcode

// https://leetcode.cn/problems/binary-tree-right-side-view
// 二叉树的右视图

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var size = len(queue)
		var right = queue[size-1] // right most node
		result = append(result, right.Val)
		for i := 0; i < size; i++ {
			var node = queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return result
}
