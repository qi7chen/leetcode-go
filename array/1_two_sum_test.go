package leetcode_array

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
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}
	for i, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		slices.Sort(got)
		assert.Equalf(t, tt.want, got, "case-%d", i+1)
	}
}

func Benchmark_twoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 6, 4, 8, 10, 14, 1, 5, 9, 12, 13}
	target := 18
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}

func Benchmark_twoSumBruteForce(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 6, 4, 8, 10, 14, 1, 5, 9, 12, 13}
	target := 18
	for i := 0; i < b.N; i++ {
		twoSumBruteForce(nums, target)
	}
}
