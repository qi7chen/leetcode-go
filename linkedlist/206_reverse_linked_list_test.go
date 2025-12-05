package leetcode_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		list := buildList(tt.input)
		reverted := reverseList(list)
		actual := fmtList(reverted)
		assert.Equalf(t, len(tt.expect), len(actual), "reverseList(%v)", tt.input)
		if len(tt.expect) > 0 {
			assert.Equalf(t, tt.expect, actual, "reverseList(%v)", tt.input)
		}
	}
}
