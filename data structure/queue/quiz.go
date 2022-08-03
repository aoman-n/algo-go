package queue

func Reverse[T any](d *Dequeue[T]) {
	newD := NewDequeue[T]()
	for n := d.DequeueBack(); n != nil; n = d.DequeueBack() {
		newD.EnqueueBack(*n)
	}

	*d = *newD
}
