package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preorderTraversal(t *testing.T) {
	//     1
	//    / \
	//   2   3
	//  /\   /\
	// 4 5  6  7
	var s = []any{1, 2, 3, 4, 5, 6, 7}
	var tree = buildTree(s...)
	var res = preorderTraversal(tree)
	assert.Equal(t, []int{1, 2, 4, 5, 3, 6, 7}, res)
}
