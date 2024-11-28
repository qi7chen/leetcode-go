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
		//{nil, nil},
		//{[]int{}, []int{}},
		{[]int{1}, []int{0}},
		{[]int{1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		out := smallerNumbersThanCurrent(tt.input)
		assert.Equal(t, len(tt.input), len(out))
		if len(tt.input) > 0 {
			assert.Equal(t, tt.expect, out)
		}
	}
}
