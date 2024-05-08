package leetcode

// https://leetcode.cn/problems/binary-tree-inorder-traversal/

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if true {
		return inorderIterative(root)
	}
	var res []int
	inorderRecursive(root, &res)
	return res
}

// 递归方式
func inorderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorderRecursive(node.Left, res)
	*res = append(*res, node.Val)
	inorderRecursive(node.Right, res)
}

// 迭代方式
func inorderIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result = make([]int, 0, 8)
	var stack = make([]*TreeNode, 0, 8)
	var node = root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]   // stack top
		stack = stack[:len(stack)-1] // pop stack

		result = append(result, node.Val)
		node = node.Right
	}
	return result
}
