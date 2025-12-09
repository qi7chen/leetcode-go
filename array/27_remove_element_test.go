package leetcode_array

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
	for _, tt := range tests {
		got := removeElement(tt.nums, tt.target)
		tt.nums = tt.nums[:got]
		assert.Equalf(t, tt.want, tt.nums, "removeElement(%v, %d)", tt.nums, tt.target)
	}
}

func Benchmark_removeElement(b *testing.B) {
	nums := []int{3, 2, 2, 3}
	target := 3
	for i := 0; i < b.N; i++ {
		removeElement(nums, target)
	}
}

func Benchmark_removeElementDoublePointer(b *testing.B) {
	nums := []int{3, 2, 2, 3}
	target := 3
	for i := 0; i < b.N; i++ {
		removeElementDoublePointer(nums, target)
	}
}
