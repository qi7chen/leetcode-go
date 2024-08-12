package leetcode_dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"climb0", 0, 1},
		{"climb1", 0, 1},
		{"climb2", 2, 2},
		{"climb3", 3, 3},
		{"climb100", 100, 1298777728820984005},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = climbStairs(tt.n)
			assert.Equalf(t, tt.want, got, "climbStairs(%v)", tt.n)
		})
	}
}
