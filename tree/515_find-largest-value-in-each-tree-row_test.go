package leetcode_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestValues(t *testing.T) {
	tests := []struct {
		input []any
		want  []int
	}{
		{[]any{1, 2, 3}, []int{1, 3}},
		{[]any{1, 3, 2, 5, 3, nil, 9}, []int{1, 3, 9}},
		{[]any{1, 2, 3, 4, 5, 6, 7}, []int{1, 3, 7}},
	}
	for i, tt := range tests {
		var name = fmt.Sprintf("case-%d", i+1)
		t.Run(name, func(t *testing.T) {
			var tree = buildTree(tt.input)
			var got = largestValues(tree)
			assert.Equal(t, tt.want, got)
		})
	}
}
