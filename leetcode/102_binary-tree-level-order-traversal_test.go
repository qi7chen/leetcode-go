package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	var tree = buildTree([]any{3, 9, 20, nil, nil, 15, 7}...)
	var result = levelOrder(tree)
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, result)
}
