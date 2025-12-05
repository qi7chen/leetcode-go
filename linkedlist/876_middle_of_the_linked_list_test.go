package leetcode_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{nil, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6}, 4},
	}
	for _, tt := range tests {
		list := buildList(tt.input)
		node := middleNode(list)
		if tt.expect == 0 {
			assert.Nil(t, node)
		} else {
			assert.NotNil(t, node)
			assert.Equalf(t, tt.expect, node.Val, "middleNode(%v)", tt.input)
		}
	}
}
