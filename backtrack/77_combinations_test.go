package leetcode_backtrack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combine(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}
	for _, tt := range tests {
		var got = combine(tt.n, tt.k)
		assert.Equalf(t, tt.want, got, "combine(%d, %d)", tt.n, tt.k)
	}
}
