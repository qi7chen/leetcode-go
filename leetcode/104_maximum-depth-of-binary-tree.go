package leetcode

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftDepth = maxDepth(root.Left)
	var rightDepth = maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
