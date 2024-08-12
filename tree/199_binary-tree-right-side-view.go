package leetcode_tree

// https://leetcode.cn/problems/binary-tree-right-side-view
// 二叉树的右视图

func rightSideView(root *TreeNode) []int {
	if false {
		rightSideView_BFS(root)
	}
	return rightSideView_DFS(root)
}

// 广度优先实现
func rightSideView_BFS(root *TreeNode) []int {
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

func _rightSideView_DFS(root *TreeNode, depth int, out *[]int) {
	if root == nil {
		return
	}
	// 当前深度第一个访问的是右子树的节点
	if len(*out) == depth {
		*out = append(*out, root.Val)
	}
	_rightSideView_DFS(root.Right, depth+1, out) // right side first
	_rightSideView_DFS(root.Left, depth+1, out)
}

// 深度优先实现（递归）
func rightSideView_DFS(root *TreeNode) []int {
	var result []int
	_rightSideView_DFS(root, 0, &result)
	return result
}
