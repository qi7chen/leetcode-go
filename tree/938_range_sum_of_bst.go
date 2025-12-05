package leetcode_tree

// https://leetcode.cn/problems/range-sum-of-bst
//
// 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

func rangeSumBST(root *TreeNode, low int, high int) int {
	if false {
		return rangeSumBST_DFS(root, low, high)
	}
	return rangeSumBST_BFS(root, low, high)
}

// 深度优先实现（递归）
func rangeSumBST_DFS(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0 // 递归终止条件
	}
	if root.Val < low {
		return rangeSumBST_DFS(root.Right, low, high) //
	} else if root.Val > high {
		return rangeSumBST_DFS(root.Left, low, high) //
	} else {
		// 区间 low <= root.Val <= high
		return root.Val + rangeSumBST_DFS(root.Left, low, high) + rangeSumBST_DFS(root.Right, low, high)
	}
}

// 广度优先实现
func rangeSumBST_BFS(root *TreeNode, low int, high int) int {
	var sum = 0
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var node = queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if node.Val < low {
			queue = append(queue, node.Right)
		} else if node.Val > high {
			queue = append(queue, node.Left)
		} else {
			sum += node.Val
			queue = append(queue, node.Left, node.Right)
		}
	}
	return sum
}
