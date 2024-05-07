package leetcode

func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorderRecursive(root, &res)
	return res
}

func postorderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	inorderRecursive(node.Left, res)
	inorderRecursive(node.Right, res)
}
