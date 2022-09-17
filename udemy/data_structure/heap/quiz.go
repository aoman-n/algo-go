package heap

type Heap[T any] struct {
	heap []T
	cmp  func(T, T) int
}

type stringCounter struct {
	val   string
	count int
}

func NewHeap[T any](cmp func(T, T) int) *Heap[T] {
	return &Heap[T]{
		heap: make([]T, 0),
		cmp:  cmp,
	}
}

func NewStringCounterHeap(strs []string) *Heap[*stringCounter] {
	m := make(map[string]int, 0)

	for _, str := range strs {
		m[str] += 1
	}

	sc := make([]*stringCounter, 0)
	for val, count := range m {
		sc = append(sc, &stringCounter{val, count})
	}

	h := BuildHeap(sc, func(a, b *stringCounter) int {
		if a.count > b.count {
			return 1
		}
		if a.count < b.count {
			return -1
		}
		return 0
	})

	return h
}

func BuildHeap[T any](list []T, cmp func(T, T) int) *Heap[T] {
	h := &Heap[T]{
		heap: list,
		cmp:  cmp,
	}

	mid := h.parent(len(h.heap) - 1)

	for i := mid; i >= 0; i-- {
		h.heapify(i)
	}

	return h
}

func (h *Heap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap[T]) left(i int) int {
	return i*2 + 1
}

func (h *Heap[T]) right(i int) int {
	return i*2 + 2
}

func (h *Heap[T]) heapify(i int) {
	l := h.left(i)
	r := h.right(i)

	biggest := i
	if l < len(h.heap) {
		if cmp := h.cmp(h.heap[l], h.heap[biggest]); cmp > 0 {
			biggest = l
		}
	}
	if r < len(h.heap) {
		if cmp := h.cmp(h.heap[r], h.heap[biggest]); cmp > 0 {
			biggest = r
		}
	}

	if biggest != i {
		h.heap[i], h.heap[biggest] = h.heap[biggest], h.heap[i]
		h.heapify(biggest)
	}
}

type Optional[T any] struct {
	Val   T
	Valid bool
}

func (h *Heap[T]) Pop() Optional[T] {
	if len(h.heap) == 0 {
		return Optional[T]{
			Valid: false,
		}
	}

	ret := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapify(0)

	return Optional[T]{
		Val:   ret,
		Valid: true,
	}
}

func TopNWithHeap(list []string, n int) []string {
	h := NewStringCounterHeap(list)

	ret := make([]string, 0, n)
	for range make([]struct{}, n) {
		data := h.Pop().Val
		ret = append(ret, data.val)
	}
	return ret
}
