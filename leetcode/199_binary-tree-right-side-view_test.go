package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		nodes  []any
		expect []int
	}{
		{[]any{1, nil, 3}, []int{1, 3}},
		{[]any{1, 2, 3, nil, 5, nil, 4}, []int{1, 3, 4}},
	}
	for i, tt := range tests {
		var tree = buildTree(tt.nodes...)
		var result = rightSideView(tree)
		assert.Equalf(t, tt.expect, result, "case-%d", i+1)
	}
}
