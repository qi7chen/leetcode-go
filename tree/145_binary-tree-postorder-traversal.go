package leetcode_tree

// 二叉树的后序遍历，左子树 -> 右子树 -> 根节点
// https://leetcode.cn/problems/binary-tree-postorder-traversal
func postorderTraversal(root *TreeNode) []int {
	if false {
		return postorderIterative(root)
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
func postorderIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result = make([]int, 0, 8)
	var stack = make([]*TreeNode, 0, 8)
	for root != nil || len(stack) > 0 {
		for root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]   // stack top
		stack = stack[:len(stack)-1] // pop stack

		// If the popped item has a right child and the right child is at top of stack,
		// then remove the right child from stack, push the root back and set root as root's right child.
		if root.Right != nil && len(stack) > 0 && root.Right == stack[len(stack)-1] {
			stack = stack[:len(stack)-1] // pop right child
			stack = append(stack, root)  // push root back
			root = root.Right
		} else {
			result = append(result, root.Val)
			root = nil
		}
	}
	return result
}

// 颜色标记法，对前序、中序、后序遍历，能够写出完全一致的代码
func postorderColorMark(root *TreeNode) []int {
	var result = make([]int, 0, 8)
	var visited = make(map[*TreeNode]struct{})
	var stack = []*TreeNode{root}
	for len(stack) > 0 {
		var node = stack[len(stack)-1] // peek stack
		stack = stack[:len(stack)-1]   // pop stack
		if node == nil {
			continue
		}
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			stack = append(stack, node)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			result = append(result, node.Val)
		}
	}
	return result
}
