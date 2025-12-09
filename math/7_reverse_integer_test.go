package leetcode_math

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{1234, 4321},
		{-1234, -4321},
		{120, 21},
		{123456789, 987654321},
		{math.MaxUint32, 0},
		{math.MaxInt32, 0},
	}
	for _, tt := range tests {
		got := reverse(tt.x)
		assert.Equal(t, tt.want, got, "reverse(%v)", tt.x)
	}
}
