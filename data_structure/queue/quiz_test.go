package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	d := NewDequeue[int]()

	d.EnqueueBack(0)
	d.EnqueueBack(1)
	d.EnqueueBack(2)
	d.EnqueueBack(3)
	d.EnqueueBack(4)
	d.EnqueueBack(5)
	Reverse(d)

	got := d.GetSlice()

	assert.Equal(t, []int{5, 4, 3, 2, 1, 0}, got)
}
