package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_invertTree(t *testing.T) {
	tests := []struct {
		tree []any
		want []any
	}{
		{[]any{4, 2, 7, 1, 3, 6, 9}, []any{4, 7, 2, 9, 6, 3, 1}},
	}
	for _, tt := range tests {
		var tree = buildTree(tt.tree)
		var inverted = invertTree(tree)
		var out = fmtTree(inverted)
		assert.Equal(t, tt.want, out)
	}
}
