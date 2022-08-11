package queue

type node[T any] struct {
	val  T
	prev *node[T]
	next *node[T]
}

type Dequeue[T any] struct {
	head *node[T]
	tail *node[T]
}

func NewDequeue[T any]() *Dequeue[T] {
	return &Dequeue[T]{}
}

func (d *Dequeue[T]) PeekFront() *T {
	if d.head == nil {
		return nil
	}
	return &d.head.val
}

func (d *Dequeue[T]) EnqueueFront(val T) {
	n := &node[T]{val: val}
	if d.head == nil {
		d.head = n
		d.tail = n
		return
	}

	n.next = d.head
	d.head.prev = n
	d.head = n
}

func (d *Dequeue[T]) DequeueFront() *T {
	if d.head == nil {
		return nil
	}

	tmp := d.head
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
		return &tmp.val
	} else {
		d.head = d.head.next
		d.head.prev = nil
		return &tmp.val
	}
}

func (d *Dequeue[T]) PeekBack() *T {
	if d.tail == nil {
		return nil
	}
	return &d.tail.val
}

func (d *Dequeue[T]) EnqueueBack(val T) {
	n := &node[T]{val: val}
	if d.head == nil {
		d.head = n
		d.tail = n
		return
	} else {
		d.tail.next = n
		n.prev = d.tail
		d.tail = n
	}
}

func (d *Dequeue[T]) DequeueBack() *T {
	if d.tail == nil {
		return nil
	}

	tmp := d.tail
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
		return &tmp.val
	} else {
		tmp := d.tail
		d.tail = d.tail.prev
		d.tail.next = nil
		return &tmp.val
	}
}

func (d *Dequeue[T]) GetSlice() []T {
	s := make([]T, 0)
	for n := d.head; n != nil; n = n.next {
		s = append(s, n.val)
	}
	return s
}
