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
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop stack

		result = append(result, root.Val)
		root = root.Right
	}
	return result
}

// https://en.wikipedia.org/wiki/Tree_traversal#Morris_in-order_traversal_using_threading
func morrisInorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result = make([]int, 0, 8)
	for root != nil {
		// 左子树为空，则打印这个节点，并向右边遍历
		if root.Left == nil {
			result = append(result, root.Val)
			root = root.Right
		} else {
			// 如果左节点不为空，就将当前节点连带右子树全部挂到左节点的最右子树下面
			var parent = root.Left
			for parent.Right != nil && parent.Right != root {
				parent = parent.Right
			}
			if parent.Right == nil {
				parent.Right = root
				root = root.Left
			} else {
				parent.Right = nil
				result = append(result, root.Val)
				root = root.Right
			}
		}
	}
	return result
}
