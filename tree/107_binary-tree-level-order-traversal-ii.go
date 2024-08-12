package leetcode_tree

import "slices"

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
// 二叉树的层序遍历 II

func levelOrderBottom(root *TreeNode) [][]int {
	if UseBFS {
		return levelOrderBottom_BFS(root)
	}
	return levelOrderBottom_DFS(root)
}

// 广度优先实现
func levelOrderBottom_BFS(root *TreeNode) [][]int {
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
		queue = queue[size:]
		// result = append([][]int{level}, result) // push front
		result = append(result, level) // push back
	}
	slices.Reverse(result)
	return result
}

// 深度优先实现（递归）
func _levelOrderBottom_DFS(root *TreeNode, depth int, out *[][]int) {
	if root == nil {
		return
	}
	if len(*out) == depth {
		*out = append(*out, []int{})
	}
	(*out)[depth] = append((*out)[depth], root.Val)
	_levelOrderBottom_DFS(root.Left, depth+1, out)
	_levelOrderBottom_DFS(root.Right, depth+1, out)
}

func levelOrderBottom_DFS(root *TreeNode) [][]int {
	var result = make([][]int, 0, 8)
	_levelOrderBottom_DFS(root, 0, &result)
	slices.Reverse(result)
	return result
}
