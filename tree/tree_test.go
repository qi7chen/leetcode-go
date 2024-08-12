package leetcode_tree

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	{
		var s []any
		for i := 1; i < 15; i++ {
			s = append(s, i)
		}
		var tree = buildTree(s)
		printTree(tree)
	}
	{
		var s = []any{1, 2, 3, nil, 5, 6, nil}
		var tree = buildTree(s)
		printTree(tree)
	}
}

func Test_fmtTree(t *testing.T) {

}
