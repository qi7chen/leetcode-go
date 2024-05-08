package leetcode

// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/

// 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	if false {
		return postorderImperative(root)
	}
	var res []int
	postorderRecursive(root, &res)
	return res
}

// 递归方式
func postorderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postorderRecursive(node.Left, res)
	postorderRecursive(node.Right, res)
	*res = append(*res, node.Val)
}

// 迭代方式
func postorderImperative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result = make([]int, 0, 8)
	var stack = make([]*TreeNode, 0, 8)
	var node = root
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			node = node.Right
		}
	}
	return result
}
