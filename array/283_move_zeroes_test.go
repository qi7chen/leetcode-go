package leetcode_array

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}
	for _, tt := range tests {
		moveZeroes(tt.input)
		var got = tt.input
		assert.True(t, slices.Equal(tt.want, got))
	}
}
