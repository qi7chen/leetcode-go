package leetcode_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	}
	for i, tt := range tests {
		var n = removeDuplicates(tt.nums)
		assert.Equalf(t, len(tt.expect), n, "case-%d", i+1)
		assert.Equalf(t, tt.expect, tt.nums[:n], "case-%d", i+1)
	}
}
