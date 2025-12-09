package leetcode_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{"0", "0", "0"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}
	for _, tt := range tests {
		out := addBinary(tt.a, tt.b)
		assert.Equalf(t, tt.want, out, "addBinary(%s, %s)", tt.a, tt.b)
	}
}
