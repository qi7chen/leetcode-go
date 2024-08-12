package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		input []any
		want  []int
	}{
		//{[]any{}, []int{}},
		//{[]any{1}, []int{1}},
		{[]any{1, nil, 2, 3}, []int{1, 3, 2}},
		{[]any{1, 2, 3, 4, 5, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7}},
		{[]any{1, nil, 0, 3}, []int{1, 3, 0}},
		{[]any{37, -34, -48, nil, -100, -100, 48, nil, nil, nil, nil, -54, nil, -71, -22, nil, nil, nil, 8}, []int{-34, -100, 37, -100, -48, -71, -54, -22, 8, 48}},
	}
	for i, tt := range tests {
		var tree = buildTree(tt.input...)
		var res = inorderTraversal(tree)
		assert.Equalf(t, tt.want, res, "case-%d", i+1)
	}
}
