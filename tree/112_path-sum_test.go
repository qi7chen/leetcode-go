package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		tree []any
		sum  int
		want bool
	}{
		{[]any{}, 0, false},
		{[]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}, 22, true},
		{[]any{1, 2}, 1, false},
	}
	for _, tt := range tests {
		var tree = buildTree(tt.tree)
		var out = hasPathSum(tree, tt.sum)
		assert.Equal(t, tt.want, out)
	}
}
