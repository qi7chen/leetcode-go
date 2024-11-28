package leetcode_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		expect []int
	}{
		{nil, nil, nil},
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{0}, []int{0, 1}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{4, 5, 6}, []int{4, 5, 6}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		list1 := buildList(tt.input1)
		list2 := buildList(tt.input2)
		merged := mergeTwoLists(list1, list2)
		actual := fmtList(merged)
		assert.Equalf(t, len(tt.expect), len(actual), "mergeTwoLists(%v,%v)", tt.input1, tt.input2)
		if len(tt.expect) > 0 {
			assert.Equalf(t, tt.expect, actual, "mergeTwoLists(%v, %v)", tt.input1, tt.input2)
		}
	}
}
