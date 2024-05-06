package leetcode

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	var s = []any{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var tree = buildTree(s...)
	printTree(tree)
}

func Test_fmtTree(t *testing.T) {

}
