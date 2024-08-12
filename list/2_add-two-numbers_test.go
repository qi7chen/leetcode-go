package leetcode_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		want   []int
	}{
		{[]int{0}, []int{0}, []int{0}},
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{9, 9, 9, 9}, []int{9, 9, 9, 9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, tt := range tests {
		l1 := buildList(tt.input1)
		l2 := buildList(tt.input2)
		out := addTwoNumbers(l1, l2)
		l3 := fmtList(out)
		assert.Equalf(t, tt.want, l3, "addTwoNumbers(%v, %v)", tt.input1, tt.input2)
	}
}
