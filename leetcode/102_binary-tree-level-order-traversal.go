package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var size = len(queue)
		var level = make([]int, 0, size)
		for i := 0; i < size; i++ {
			var node = queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			level = append(level, node.Val)
		}
		result = append(result, level)
	}
	return result
}
