package leetcode_list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		input []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		var name = fmt.Sprintf("removeNthFromEnd(%d)", tt.n)
		t.Run(name, func(t *testing.T) {
			var list = buildList(tt.input)
			var head = removeNthFromEnd(list, tt.n)
			var nodes = fmtList(head)
			assert.Equal(t, tt.want, nodes)
		})
	}
}
