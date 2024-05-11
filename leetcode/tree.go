package leetcode

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxTreeDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var leftDepth = maxTreeDepth(node.Left)
	var rightDepth = maxTreeDepth(node.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func createTreeNodeBy(v any) *TreeNode {
	val, err := toInt(v)
	if err != nil {
		return nil
	}
	return &TreeNode{Val: val}
}

func buildTree(s ...any) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	var root = createTreeNodeBy(s[0])
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	for i := 1; i < len(s); {
		var parent = queue[0]
		queue = queue[1:] // pop queue
		if i < len(s) {
			var node = createTreeNodeBy(s[i]) // left child
			i++
			if node != nil {
				parent.Left = node
				queue = append(queue, node)
			}
		}
		if i < len(s) {
			var node = createTreeNodeBy(s[i]) // right child
			i++
			if node != nil {
				parent.Right = node
				queue = append(queue, node)
			}
		}
	}
	return root
}

func fmtTree(root *TreeNode) []any {
	if root == nil {
		return nil
	}
	var result []any
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var node = queue[0]
		queue = queue[1:] // pop left
		if node == nil {
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	return result
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	var maxDepth = maxTreeDepth(root)
	var queue = []*TreeNode{root}
	var depth = 1
	for len(queue) > 0 {
		var size = len(queue)
		for i := 0; i < size; i++ {
			var node = queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			printNode(node, maxDepth, depth)
		}
		depth++
		queue = queue[size:] // pop queue
		fmt.Println()
	}
}

func printNode(node *TreeNode, maxDepth, depth int) {
	var padding = strings.Repeat(" ", maxDepth-depth+1)
	if node == nil {
		fmt.Printf("%snull", padding)
	} else {
		fmt.Printf("%s%d", padding, node.Val)
	}
}
