package leetcode_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, tt := range tests {
		out := isValid(tt.input)
		assert.Equalf(t, out, tt.want, "input: %v", tt.input)
	}
}
