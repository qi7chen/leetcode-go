package leetcode_tree

// 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var right = invertTree(root.Right)
	var left = invertTree(root.Left)
	root.Left = right
	root.Right = left
	return root
}
