package leetcode_array

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		input []int
		want  [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		got := threeSum(tt.input)
		assert.Equalf(t, len(tt.want), len(got), "threeSum(%v)", tt.input)
		for i := 0; i < len(tt.want) && i < len(got); i++ {
			assert.Truef(t, slices.Equal(tt.want[i], got[i]), "threeSum(%v)", tt.input)
		}
	}
}
