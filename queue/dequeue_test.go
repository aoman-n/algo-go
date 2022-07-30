package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeue(t *testing.T) {

	t.Run("queue動作", func(t *testing.T) {
		q := NewDequeue[int]()

		// queue動作
		q.EnqueueBack(1)
		q.EnqueueBack(2)
		q.EnqueueBack(3)
		assert.Equal(t, 1, *q.DequeueFront())
		assert.Equal(t, 2, *q.DequeueFront())
		assert.Equal(t, 3, *q.DequeueFront())
		if got := q.DequeueFront(); got != nil {
			t.Errorf("want nil, but got %v", got)
		}
	})

	t.Run("逆queue動作", func(t *testing.T) {
		q := NewDequeue[int]()

		q.EnqueueFront(3)
		q.EnqueueFront(2)
		q.EnqueueFront(1)
		assert.Equal(t, 3, *q.DequeueBack())
		assert.Equal(t, 2, *q.DequeueBack()) //
		assert.Equal(t, 1, *q.DequeueBack())
		if got := q.DequeueBack(); got != nil {
			t.Errorf("want nil, but got %v", got)
		}
	})

	t.Run("両方混在", func(t *testing.T) {
		q := NewDequeue[int]()

		q.EnqueueBack(1)
		q.EnqueueBack(2)
		q.EnqueueBack(3)
		q.EnqueueFront(0)
		q.EnqueueFront(-1)

		assert.Equal(t, -1, *q.DequeueFront())
		assert.Equal(t, 3, *q.DequeueBack())

		q.EnqueueBack(3)
		assert.Equal(t, 3, *q.DequeueBack())
		assert.Equal(t, 2, *q.DequeueBack())
		assert.Equal(t, 1, *q.DequeueBack())
		assert.Equal(t, 0, *q.DequeueFront())
	})
}
