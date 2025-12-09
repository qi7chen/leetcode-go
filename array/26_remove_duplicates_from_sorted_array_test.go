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
	for _, tt := range tests {
		var got = removeDuplicates(tt.nums)
		assert.Equalf(t, len(tt.expect), got, "removeDuplicates(%v)", tt.nums)
		for j := 0; j < got; j++ {
			assert.Equalf(t, tt.expect[j], tt.nums[j], "removeDuplicates(%v)", tt.nums)
		}
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	for i := 0; i < b.N; i++ {
		removeDuplicates(nums)
	}
}
