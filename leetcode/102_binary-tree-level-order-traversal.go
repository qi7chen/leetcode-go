package leetcode

// https://leetcode.cn/problems/binary-tree-level-order-traversal
// 二叉树层序遍历

func levelOrder(root *TreeNode) [][]int {
	if UseBFS {
		return levelOrder_BFS(root)
	}
	return levelOrder_DFS(root)
}

// 广度优先实现
func levelOrder_BFS(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result = make([][]int, 0, 8)
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
		result = append(result, level) // push back
	}
	return result
}

// 深度优先实现（递归）
func _levelOrder_DFS(root *TreeNode, depth int, out *[][]int) {
	if root == nil {
		return
	}
	if len(*out) == depth {
		*out = append(*out, []int{})
	}
	(*out)[depth] = append((*out)[depth], root.Val)
	_levelOrder_DFS(root.Left, depth+1, out)
	_levelOrder_DFS(root.Right, depth+1, out)
}

func levelOrder_DFS(root *TreeNode) [][]int {
	var result = make([][]int, 0, 8)
	_levelOrder_DFS(root, 0, &result)
	return result
}
