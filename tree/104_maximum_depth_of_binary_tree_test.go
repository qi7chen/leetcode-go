package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		input []any
		want  int
	}{
		{[]any{3, 9, 20, nil, nil, 15, 7}, 3},
	}
	for i, tt := range tests {
		var tree = buildTree(tt.input)
		var depth = maxDepth(tree)
		assert.Equalf(t, tt.want, depth, "case %d maxDepth() %d != %d", i+1, depth, tt.want)
	}
}
