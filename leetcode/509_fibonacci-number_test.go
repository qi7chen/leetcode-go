package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"fib0", 0, 0},
		{"fib1", 1, 1},
		{"fib2", 100, 3736710778780434371},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, fib(tt.n), "fib(%v)", tt.n)
		})
	}
}
