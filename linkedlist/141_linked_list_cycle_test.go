package leetcode_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setListCycle(head *ListNode, n int) {
	var tail = getListTail(head)
	var node = head
	for i := 0; i < n; i++ {
		node = node.Next
	}
	tail.Next = node
}

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		input    []int
		setCycle int
		hasCycle bool
	}{
		{nil, 0, false},
		{[]int{}, 0, false},
		{[]int{1}, 0, false},
		{[]int{1, 2, 3}, 1, true},
	}
	for _, tt := range tests {
		list := buildList(tt.input)
		if tt.setCycle > 0 {
			setListCycle(list, tt.setCycle)
		}
		output := hasCycle(list)
		assert.Equalf(t, tt.hasCycle, output, "hasCycle(%v)", tt.input)
	}
}
