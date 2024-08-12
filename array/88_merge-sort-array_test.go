package leetcode_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		num1 []int
		m    int
		num2 []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		merge(tt.num1, tt.m, tt.num2, tt.n)
		assert.Equal(t, tt.want, tt.num1)
	}
}
