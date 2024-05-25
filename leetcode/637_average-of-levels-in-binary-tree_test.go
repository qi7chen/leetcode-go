package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_averageOfLevels(t *testing.T) {
	tests := []struct {
		nodes  []any
		expect []float64
	}{
		{[]any{3, 9, 20, nil, nil, 15, 7}, []float64{3, 14.5, 11}},
		{[]any{3, 9, 20, 15, 7}, []float64{3, 14.5, 11}},
	}
	for i, tt := range tests {
		var tree = buildTree(tt.nodes...)
		var result = averageOfLevels(tree)
		assert.Equalf(t, tt.expect, result, "case-%d", i+1)
	}
}
