package leetcode_tree

// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
// N 叉树的层序遍历

func nAryLevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var queue = []*Node{root}
	for len(queue) > 0 {
		var size = len(queue)
		var level = make([]int, 0, size)
		for i := 0; i < size; i++ {
			var node = queue[i]
			for _, child := range node.Children {
				if child != nil {
					queue = append(queue, child)
				}
			}
			level = append(level, node.Val)
		}
		queue = queue[size:]
		result = append(result, level)
	}
	return result
}
