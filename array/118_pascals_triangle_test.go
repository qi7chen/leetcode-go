package leetcode_array

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	var tests = []struct {
		input int
		want  [][]int
	}{
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
		{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}
	for _, tt := range tests {
		var got = generate(tt.input)
		assert.Equal(t, len(tt.want), len(got))
		for i := range tt.want {
			assert.True(t, slices.Equal(got[i], tt.want[i]))
		}
	}
}
