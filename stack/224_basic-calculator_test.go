package leetcode_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate(t *testing.T) {

	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"  30", 30},
		{"1+2", 3},
		{"1 + 1", 2},
		{"2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
	}
	for _, tt := range tests {
		out := calculate(tt.input)
		assert.Equalf(t, out, tt.want, "input: %v", tt.input)
	}
}
