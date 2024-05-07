package leetcode

func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorderRecursive(root, &res)
	return res
}

func inorderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorderRecursive(node.Left, res)
	*res = append(*res, node.Val)
	inorderRecursive(node.Right, res)
}

func inorderImperative(node *TreeNode) []int {
	// TODO:
	return nil
}
