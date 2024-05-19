package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{}, 9, []int{}},
		{[]int{3, 2, 2, 3}, 3, []int{2, 2}},
		{[]int{1, 2, 3, 4}, 1, []int{2, 3, 4}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3}},
	}
	for i, tt := range tests {
		got := removeElement(tt.nums, tt.target)
		tt.nums = tt.nums[:got]
		assert.Equalf(t, tt.want, tt.nums, "case-%d", i+1)
	}
}
