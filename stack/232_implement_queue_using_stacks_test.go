package leetcode_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueue_Example(t *testing.T) {
	{
		var q = Constructor()
		assert.True(t, q.Empty())
		q.Push(111)
		q.Push(222)
		assert.False(t, q.Empty())
		var val = q.Peek()
		assert.Equal(t, 111, val)
		assert.False(t, q.Empty())
		val = q.Pop()
		assert.Equal(t, 111, val)
		val = q.Pop()
		assert.Equal(t, 222, val)
		assert.True(t, q.Empty())
	}
}
