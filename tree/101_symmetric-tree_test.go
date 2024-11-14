package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		tree []any
		want bool
	}{
		{[]any{1, 2, 2, 3, 4, 4, 3}, true},
	}
	for _, tt := range tests {
		var tree = buildTree(tt.tree)
		var out = isSymmetric(tree)
		assert.Equal(t, tt.want, out)
	}
}
