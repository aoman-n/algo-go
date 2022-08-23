package heap

type Heap struct {
	data []int
	size int
}

func newHeap(list []int) *Heap {
	return &Heap{
		data: list,
		size: len(list),
	}
}

func BuildMaxHeap(list []int) *Heap {
	newList := make([]int, len(list))
	copy(newList, list)
	heap := newHeap(newList)

	startIndex := heap.parent(len(newList) - 1)
	for i := startIndex; i >= 0; i-- {
		heap.maxHeapify(i, len(heap.data)-1)
	}

	return heap
}

func (h Heap) List() []int {
	return h.data
}

func (h *Heap) maxHeapify(i, end int) {
	l := h.left(i)
	r := h.right(i)

	biggest := i
	if l < end && h.data[l] > h.data[biggest] {
		biggest = l
	}
	if r < end && h.data[r] > h.data[biggest] {
		biggest = r
	}

	if i != biggest {
		h.swap(i, biggest)
		h.maxHeapify(biggest, end)
	}
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) left(i int) int {
	return i*2 + 1
}

func (h *Heap) right(i int) int {
	return i*2 + 2
}

func (h *Heap) parent(i int) int {
	return i/2 - 1
}

func (h *Heap) Sort() {
	for i := len(h.data) - 1; i >= 0; i-- {
		h.swap(0, i)
		h.maxHeapify(0, i)
	}
}

func Sort(list []int) []int {
	maxHeap := BuildMaxHeap(list)
	maxHeap.Sort()
	return maxHeap.data
}
