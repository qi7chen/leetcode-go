package leetcode_tree

// https://leetcode.cn/problems/binary-tree-preorder-traversal/

// 二叉树的前序遍历：根节点 -> 左子树 -> 右子树
func preorderTraversal(root *TreeNode) []int {
	if UseIterativeTraversal {
		return preorderIterative(root)
	}
	return preorderColorMark(root)
}

// 递归方式
func preorderRecursive(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preorderRecursive(node.Left, res)
	preorderRecursive(node.Right, res)
}

// 迭代方式
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

		// 先加入右子树，再加入左子树，这样出栈时才能先处理左子树
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

// 颜色标记法，对前序、中序、后序遍历，能够写出完全一致的代码
func preorderColorMark(root *TreeNode) []int {
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
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			visited[node] = struct{}{}
			stack = append(stack, node)
		} else {
			result = append(result, node.Val)
		}
	}
	return result
}
