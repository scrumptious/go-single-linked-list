package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	ll1 := &LinkedList{}
	expected := []int{4, 5, 6}

	ll1.Append(expected...)
	assert.Equal(t, expected[0], ll1.head.val)
	assert.Equal(t, expected[1], ll1.head.next.val)
	assert.Equal(t, expected[2], ll1.head.next.next.val)

}
