package leetcode_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"", "", true},
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a#c", "b", false},
	}
	for i, tt := range tests {
		out := backspaceCompare(tt.s, tt.t)
		assert.Equalf(t, tt.want, out, "case-%d", i)
	}
}
