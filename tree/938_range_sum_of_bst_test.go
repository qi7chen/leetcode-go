package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rangeSumBST(t *testing.T) {
	tests := []struct {
		nodes  []any
		input1 int
		input2 int
		expect int
	}{
		{[]any{10, 5, 15, 3, 7, nil, 18}, 7, 15, 32},
	}
	for i, tt := range tests {
		var tree = buildTree(tt.nodes)
		var result = rangeSumBST(tree, tt.input1, tt.input2)
		assert.Equalf(t, tt.expect, result, "case-%d", i+1)
	}
}
