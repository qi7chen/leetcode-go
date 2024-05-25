package leetcode

// https://leetcode.cn/problems/binary-tree-preorder-traversal/

// 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	if UseIterativeTraversal {
		return preorderIterative(root)
	}
	var res = make([]int, 0, 8)
	preorderRecursive(root, &res)
	return res
}

// 递归方式的前序遍历
func preorderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preorderRecursive(node.Left, res)
	preorderRecursive(node.Right, res)
}

// 迭代方式的前序遍历
func preorderIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result = make([]int, 0, 8)
	var stack = make([]*TreeNode, 0, 8)
	stack = append(stack, root)
	for len(stack) > 0 {
		var node = stack[len(stack)-1] // stack top
		stack = stack[:len(stack)-1]   // pop stack

		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}
