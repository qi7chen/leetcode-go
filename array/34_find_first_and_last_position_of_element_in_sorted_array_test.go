package leetcode_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 15, []int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 5, []int{0, 0}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	}
	for _, tt := range tests {
		got := searchRange(tt.nums, tt.target)
		assert.Equalf(t, tt.want, got, "searchRange(%v, %d)", tt.nums, tt.target)
	}
}
