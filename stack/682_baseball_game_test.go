package leetcode_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calPoints(t *testing.T) {
	tests := []struct {
		ops  []string
		want int
	}{
		{nil, 0},
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
	}
	for _, tt := range tests {
		out := calPoints(tt.ops)
		assert.Equalf(t, tt.want, out, "input: %v", tt.ops)
	}
}
