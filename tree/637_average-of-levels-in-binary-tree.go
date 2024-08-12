package leetcode_tree

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/
// 二叉树的层平均值

func averageOfLevels(root *TreeNode) []float64 {
	if UseBFSOverDFS {
		return averageOfLevels_BFS(root)
	}
	return averageOfLevels_DFS(root)
}

// 广度优先实现
func averageOfLevels_BFS(root *TreeNode) []float64 {
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

// 深度优先实现（递归）
func _averageOfLevels_DFS(root *TreeNode, depth int, outCount, outSum *[]int) {
	if root == nil {
		return
	}
	if len(*outCount) == depth {
		*outCount = append(*outCount, 0)
		*outSum = append(*outSum, 0)
	}
	(*outCount)[depth] += 1
	(*outSum)[depth] += root.Val
	_averageOfLevels_DFS(root.Left, depth+1, outCount, outSum)
	_averageOfLevels_DFS(root.Right, depth+1, outCount, outSum)
}

// 深度优先实现
func averageOfLevels_DFS(root *TreeNode) []float64 {
	var counts = make([]int, 0, 8)
	var sums = make([]int, 0, 8)
	_averageOfLevels_DFS(root, 0, &counts, &sums)
	var result = make([]float64, len(counts))
	for i := 0; i < len(counts); i++ {
		result[i] = float64(sums[i]) / float64(counts[i])
	}
	return result
}
