package leetcode_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	tests := []struct {
		input1 []any
		input2 []any
		want   bool
	}{
		{[]any{1, 2, 3}, []any{1, 2, 3}, true},
		{[]any{1, 2}, []any{1, nil, 2}, false},
		{[]any{1, 2, 1}, []any{1, 1, 2}, false},
	}
	for _, tt := range tests {
		var tree1 = buildTree(tt.input1...)
		var tree2 = buildTree(tt.input2...)
		assert.Equal(t, tt.want, isSameTree(tree1, tree2))
	}
}
