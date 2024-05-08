package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_postorderTraversal(t *testing.T) {
	//     1
	//    / \
	//   2   3
	//  /\   /\
	// 4 5  6  7
	var s = []any{1, 2, 3, 4, 5, 6, 7}
	var tree = buildTree(s...)
	var res = postorderTraversal(tree)
	assert.Equal(t, []int{4, 5, 2, 6, 7, 3, 1}, res)
}
