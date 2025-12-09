package leetcode_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{0}},
		{[]int{1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4}},
		{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
		{[]int{7, 7, 7, 7}, []int{0, 0, 0, 0}},
	}
	for _, tt := range tests {
		out := smallerNumbersThanCurrent(tt.input)
		assert.Equalf(t, len(tt.input), len(out), "smallerNumbersThanCurrent(%v)", tt.input)
		if len(tt.input) > 0 {
			assert.Equalf(t, tt.expect, out, "smallerNumbersThanCurrent(%v)", tt.input)
		}
	}
}
