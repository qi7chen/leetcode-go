package leetcode_math

import (
	"fmt"
	"math"
	"testing"
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
		var name = fmt.Sprintf("reverse(%v)", tt.want)
		t.Run(name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
