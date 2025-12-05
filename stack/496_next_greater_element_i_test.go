package leetcode_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextGreaterElement(t *testing.T) {

	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		out := nextGreaterElement(tt.nums1, tt.nums2)
		assert.Equalf(t, out, tt.want, "input: %v", tt.nums1)
	}
}
