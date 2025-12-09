package leetcode_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{}, 9, 0},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 3, 6}, 3, 1},
	}
	for _, tt := range tests {
		got := searchInsert(tt.nums, tt.target)
		assert.Equalf(t, tt.want, got, "searchInsert(%v, %d)", tt.nums, tt.target)
	}
}
