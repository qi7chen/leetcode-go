package leetcode

// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
// 在每个树行中找最大值

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var size = len(queue)
		var max = queue[0].Val
		for i := 0; i < size; i++ {
			var node = queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Val > max {
				max = node.Val
			}
		}
		queue = queue[size:]
		result = append(result, max)
	}
	return result
}
