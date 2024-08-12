package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrderBottom(t *testing.T) {
	tests := []struct {
		nodes  []any
		expect [][]int
	}{
		{[]any{1}, [][]int{{1}}},
		{[]any{3, 9, 20, nil, nil, 15, 7}, [][]int{{15, 7}, {9, 20}, {3}}},
	}
	for i, tt := range tests {
		var tree = buildTree(tt.nodes...)
		var result = levelOrderBottom(tree)
		assert.Equalf(t, tt.expect, result, "case-%d", i+1)
	}
}
