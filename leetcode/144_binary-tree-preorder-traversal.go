package leetcode

func preorderTraversal(root *TreeNode) []int {
	var res []int
	preorderRecursive(root, &res)
	return nil
}

func preorderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	inorderRecursive(node.Left, res)
	inorderRecursive(node.Right, res)
}
