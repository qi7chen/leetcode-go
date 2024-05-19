package leetcode

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/
// 二叉树的层平均值

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var size = len(queue)
		var sum float64
		for i := 0; i < size; i++ {
			var node = queue[i]
			sum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		result = append(result, sum/float64(size))
	}
	return result
}
