package leetcode_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getDecimalValue(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{nil, 0},
		{[]int{0}, 0},
		{[]int{0, 0}, 0},
		{[]int{1}, 1},
		{[]int{1, 0, 1}, 5},
		{[]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}, 18880},
	}
	for _, tt := range tests {
		list := buildList(tt.input)
		result := getDecimalValue(list)
		assert.Equalf(t, tt.expect, result, "getDecimalValue(%v)", tt.input)
	}
}
