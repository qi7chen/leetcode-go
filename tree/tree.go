package leetcode_tree

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ColorWhite = 0 // 未访问过的节点
	ColorGray  = 1 // 访问过的节点
)

const (
	TraversalIterative = true  // 使用迭代遍历，而不是递归
	UseBFSOverDFS      = false // 使用广度优先搜索（BFS），不是深度优先搜索（DFS）
)

// Node N叉树
type Node struct {
	Val      int
	Children []*Node
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用node.Val的高位标记颜色
func markColor(node *TreeNode) {
	if node.Val >= 0 {
		node.Val += math.MaxUint32
	} else {
		node.Val -= math.MaxUint32
	}
}

func unmarkColor(node *TreeNode) {
	if node.Val >= math.MaxUint32 {
		node.Val -= math.MaxUint32
	} else {
		node.Val += math.MaxUint32
	}
}

func getColor(node *TreeNode) int8 {
	if node.Val >= math.MaxUint32 || node.Val <= math.MinInt32 {
		return ColorGray
	}
	return ColorWhite
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

func toInt(v any) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	}
	return strconv.Atoi(fmt.Sprintf("%v", v))
}

func createTreeNodeBy(v any) *TreeNode {
	val, err := toInt(v)
	if err != nil {
		return nil
	}
	return &TreeNode{Val: val}
}

func buildTree(s []any) *TreeNode {
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
