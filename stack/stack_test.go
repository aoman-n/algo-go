package stack

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	buf := &bytes.Buffer{}
	s.Print(buf)
	assert.Equal(t, "4 3 2 1\n", buf.String())
	assert.Equal(t, 4, s.Len())

	assert.Equal(t, 4, s.Pop().Val)
	assert.Equal(t, 3, s.Len())

	assert.Equal(t, 3, s.Pop().Val)
	assert.Equal(t, 2, s.Len())

	assert.Equal(t, 2, s.Pop().Val)
	assert.Equal(t, 1, s.Len())

	assert.Equal(t, 1, s.Pop().Val)
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, false, s.Pop().Valid)
}
