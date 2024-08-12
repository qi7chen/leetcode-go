package leetcode_tree

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if UseBFSOverDFS {
		return maxDepth_BFS(root)
	}
	return maxDepth_DFS(root)
}

// 深度优先算法（递归）
func maxDepth_DFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftDepth = maxDepth_DFS(root.Left)
	var rightDepth = maxDepth_DFS(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// 广度优先算法
func maxDepth_BFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var size = len(queue)
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
		depth++
	}
	return depth
}
