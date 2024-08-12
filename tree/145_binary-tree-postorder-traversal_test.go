package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_postorderTraversal(t *testing.T) {
	tests := []struct {
		input []any
		want  []int
	}{
		{[]any{}, nil},
		{[]any{1}, []int{1}},
		{[]any{1, nil, 2, 3}, []int{3, 2, 1}},
		{[]any{1, 2, 3, 4, 5, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}},
		{[]any{1, nil, 0, 3}, []int{3, 0, 1}},
		{[]any{37, -34, -48, nil, -100, -100, 48, nil, nil, nil, nil, -54, nil, -71, -22, nil, nil, nil, 8}, []int{-100, -34, -100, -71, 8, -22, -54, 48, -48, 37}},
	}
	for i, tt := range tests {
		var tree = buildTree(tt.input...)
		var res = postorderTraversal(tree)
		assert.Equalf(t, tt.want, res, "case-%d", i+1)
	}
}
