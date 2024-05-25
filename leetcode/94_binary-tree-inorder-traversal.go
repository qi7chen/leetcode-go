package leetcode

// https://leetcode.cn/problems/binary-tree-inorder-traversal/

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if UseIterativeTraversal {
		return inorderIterative(root)
	}
	var res = make([]int, 0, 8)
	inorderRecursive(root, &res)
	return res
}

// 递归方式的中序遍历
func inorderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorderRecursive(node.Left, res)
	*res = append(*res, node.Val)
	inorderRecursive(node.Right, res)
}

// 迭代方式的中序遍历
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

// 还有一个Morris Traversal的方法，可以实现O(1)空间复杂度的中序遍历，但是不太好理解，这里不写了
// https://en.wikipedia.org/wiki/Tree_traversal#Morris_in-order_traversal_using_threading
