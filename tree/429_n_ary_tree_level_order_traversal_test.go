package leetcode_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nAryLevelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nAryLevelOrder(tt.args.root), "nAryLevelOrder(%v)", tt.args.root)
		})
	}
}
