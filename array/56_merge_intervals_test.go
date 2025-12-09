package leetcode_array

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeInterval(t *testing.T) {
	tests := []struct {
		input [][]int
		want  [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
		{[][]int{{4, 7}, {1, 4}}, [][]int{{1, 7}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {0, 0}}, [][]int{{0, 0}, {1, 4}}},
		{[][]int{{1, 4}, {0, 1}}, [][]int{{0, 4}}},
	}
	for _, tt := range tests {
		got := mergeInterval(tt.input)
		assert.Equal(t, len(tt.want), len(got))
		for i := 0; i < len(tt.want) && i < len(got); i++ {
			var equal = slices.Equal(tt.want[i], got[i])
			assert.Truef(t, equal, "mergeInterval(%v)[%d]", tt.input, i)
		}
	}
}
