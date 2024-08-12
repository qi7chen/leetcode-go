package leetcode_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildList(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var head = buildList(a)
	assert.NotNil(t, head)
	var i = 0
	var node = head
	for ; node != nil; node = node.Next {
		assert.Equal(t, a[i], node.Val)
		i++
	}
	assert.Nil(t, node)
	assert.Equal(t, len(a), i)
}

func Test_fmtList(t *testing.T) {
	var prev, head *ListNode
	for i := 1; i < 10; i++ {
		var node = &ListNode{Val: i}
		if prev != nil {
			prev.Next = node
		} else {
			head = node
		}
		prev = node
	}
	var a = fmtList(head)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, a)
}
