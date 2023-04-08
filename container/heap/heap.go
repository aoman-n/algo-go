package heap

type heap[T any] struct {
	list []T
	//	cmpFn(x, y)
	//
	// x　<　y　だったら、-1
	// x　=　y　だったら、 0
	// x　>　y　だったら、 1
	cmp func(x, y T) int
}

func newHeap[T any](list []T, cmp func(x, y T) int) *heap[T] {
	return &heap[T]{
		list: list,
		cmp:  cmp,
	}
}

func (h *heap[T]) accessible(i int) bool {
	return i < len(h.list)
}

func (h *heap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (h *heap[T]) left(i int) int {
	return (i * 2) + 1
}

func (h *heap[T]) right(i int) int {
	return (i * 2) + 2
}

// 引数で渡されたindexから下をmaxHeapifyする
func (h *heap[T]) maxHeapify(i int) {
	biggest := i
	l := h.left(i)
	r := h.right(i)

	if h.accessible(l) && h.cmp(h.list[l], h.list[biggest]) > 0 {
		biggest = l
	}
	if h.accessible(r) && h.cmp(h.list[r], h.list[biggest]) > 0 {
		biggest = r
	}

	if i != biggest {
		tmp := h.list[i]
		h.list[i] = h.list[biggest]
		h.list[biggest] = tmp

		h.maxHeapify(biggest)
	}
}

func (h heap[T]) List() []T {
	return h.list
}

func BuildMaxHeap[T any](list []T, cmp func(x, y T) int) *heap[T] {
	heap := newHeap(list, cmp)

	i := heap.parent(len(list) - 1)
	for i >= 0 {
		heap.maxHeapify(i)
		i--
	}

	return heap
}

// 引数で渡されたindexから下をminHeapifyする
func (h *heap[T]) minHeapify(i int) {
	panic("no implemented")
}

func (h *heap[T]) swap(i, j int) {
	tmp := h.list[i]
	h.list[i] = h.list[j]
	h.list[j] = tmp
}

func (h *heap[T]) push(data T) {
	h.list = append(h.list, data)
}

// 最大値or最小値を削除して返す
func (h *heap[T]) Pop() T {
	if len(h.list) <= 0 {
		panic("no list") // FIXME
	}

	ret := h.list[0]

	lastIndex := len(h.list) - 1
	h.swap(0, lastIndex)

	h.list = h.list[:lastIndex]
	h.maxHeapify(0)

	return ret
}

// 値を挿入してheapifyする
func (h *heap[T]) Insert(data T) {
	h.push(data)

	i := len(h.list) - 1
	parent := h.parent(i)

	for parent >= 0 && h.cmp(h.list[i], h.list[parent]) > 0 {
		h.swap(i, parent)
		i = parent
		parent = h.parent(i)
	}
}

// 最大値or最小値を返す
func (h *heap[T]) Top() T {
	if len(h.list) <= 0 {
		panic("no list") // FIXME
	}

	return h.list[0]
}

// TODO:
// - 最小ヒープ
// - ヒープソート
