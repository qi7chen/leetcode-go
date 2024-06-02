package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preorderTraversal(t *testing.T) {
	tests := []struct {
		input []any
		want  []int
	}{
		{[]any{}, []int{}},
		{[]any{1}, []int{1}},
		{[]any{1, nil, 2, 3}, []int{1, 2, 3}},
		{[]any{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 4, 5, 3, 6, 7}},
		{[]any{1, nil, 0, 3}, []int{1, 0, 3}},
		{[]any{37, -34, -48, nil, -100, -100, 48, nil, nil, nil, nil, -54, nil, -71, -22, nil, nil, nil, 8}, []int{37, -34, -100, -48, -100, 48, -54, -71, -22, 8}},
	}
	for i, tt := range tests {
		var tree = buildTree(tt.input...)
		var res = preorderTraversal(tree)
		assert.Equalf(t, tt.want, res, "case-%d", i+1)
	}
}
