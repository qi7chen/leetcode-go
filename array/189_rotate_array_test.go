package leetcode_array

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	var tests = []struct {
		input []int
		k     int
		want  []int
	}{
		//{},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, []int{7, 1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, []int{6, 7, 1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}
	for _, tt := range tests {
		rotate(tt.input, tt.k)
		var got = tt.input
		assert.Truef(t, slices.Equal(tt.want, got), "rotate(%v, %d)", tt.input, tt.k)
	}
}
