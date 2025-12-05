package leetcode_tree

// 二叉树的中序遍历，左子树 -> 根节点 -> 右子树
// https://leetcode.cn/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	if TraversalIterative {
		return inorderIterative(root)
	}
	var res = make([]int, 0, 8)
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
			root = root.Left // 最左侧节点
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop stack

		result = append(result, root.Val)
		root = root.Right // 处理右子树
	}
	return result
}

// 颜色标记法，对前序、中序、后序遍历，能够写出完全一致的代码
func inorderColorMark(root *TreeNode) []int {
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
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			visited[node] = struct{}{}
			stack = append(stack, node)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			result = append(result, node.Val)
		}
	}
	return result
}

// 还有一个Morris Traversal的方法，可以实现O(1)空间复杂度的中序遍历，但是不太好理解，这里不写了
// https://en.wikipedia.org/wiki/Tree_traversal#Morris_in-order_traversal_using_threading
