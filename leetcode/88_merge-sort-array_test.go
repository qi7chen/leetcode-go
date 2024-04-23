package leetcode

import (
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		x []int
		m int
		y []int
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.x, tt.args.m, tt.args.y, tt.args.n)
		})
	}
}
