package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{}, 9, []int{-1, -1}},
		{[]int{2, 7, 11, 15}, 100, []int{-1, -1}},
		{[]int{9}, 9, []int{-1, -1}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 1, 8}, 9, []int{0, 1}},
		{[]int{-2, -7, 11, 15}, 4, []int{1, 2}},
		{[]int{-2, -7, 11, 15}, 0, []int{-1, -1}},
	}
	for i, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		slices.Sort(got)
		assert.Equalf(t, tt.want, got, "case-%d", i+1)
	}
}
