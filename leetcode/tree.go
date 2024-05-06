package leetcode

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(s ...any) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	rootVal, err := toInt(s[0])
	if err != nil {
		return nil
	}
	var root = &TreeNode{Val: rootVal}
	var queue = []*TreeNode{root}
	for i := 1; i < len(s); {
		var node = queue[0]
		queue = queue[1:] // pop left
		if i < len(s) {
			nodeVal, er := toInt(s[i])
			i++
			if er == nil {
				node.Left = &TreeNode{Val: nodeVal}
				queue = append(queue, node.Left)
			}
		}
		if i < len(s) {
			nodeVal, er := toInt(s[i])
			i++
			if er == nil {
				node.Right = &TreeNode{Val: nodeVal}
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

func fmtTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
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
	printTree(root.Left)
	fmt.Printf("%v ", root.Val)
	printTree(root.Right)
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
	return 0, fmt.Errorf("invalid type %T", v)
}
