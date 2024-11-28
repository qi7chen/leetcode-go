package leetcode_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		list := buildList(tt.input)
		removed := deleteDuplicates(list)
		actual := fmtList(removed)
		assert.Equalf(t, len(tt.expect), len(actual), "deleteDuplicates(%v)", tt.input)
		if len(tt.expect) > 0 {
			assert.Equalf(t, tt.expect, actual, "deleteDuplicates(%v)", tt.input)
		}
	}
}
