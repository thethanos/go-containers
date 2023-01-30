package containers

type сomparator[T any] func(left, right T) bool

type Heap[T any] struct {
	data []T
	cmp  сomparator[T]
}

// Accepts type сomparator[T any] func(left, right T) bool.
// For max heap comparator returns left > right.
// For min heap comparator returns left < right.
func NewHeap[T any](cmp сomparator[T]) Heap[T] {
	return Heap[T]{cmp: cmp}
}

func (h *Heap[T]) Push(value T) {

	h.data = append(h.data, value)

	for index := len(h.data) - 1; index > 0; index = h.parent(index) {

		if h.cmp(h.data[index], h.data[h.parent(index)]) {
			h.swap(index, h.parent(index))
		}
	}
}

// item T - item to delete
// at least the field to be checked in cmp  must be filled
//
// cmp comparator[T] - comparator
// if nil - default comparator is used
func (h *Heap[T]) Extract(item T, cmp сomparator[T]) T {
	if h.Empty() {
		panic("cannot extract value, heap is empty")
	}

	if cmp == nil {
		cmp = h.cmp
	}

	for index, element := range h.data {
		if !cmp(item, element) && !cmp(element, item) {
			h.remove(index)
			h.heapify(0)
			return element
		}
	}
	panic("cannot extract value, no such element in heap")
}

// item T - item to check
// at least the field to be checked in cmp must be filled
//
// cmp comparator[T] - comparator
// if nil - default comparator is used
func (h Heap[T]) Contains(item T, cmp сomparator[T]) bool {
	if cmp == nil {
		cmp = h.cmp
	}

	for _, element := range h.data {
		if !cmp(item, element) && !cmp(element, item) {
			return true
		}
	}
	return false
}

// Adds slice data to the heap.
func (h *Heap[T]) PushSlice(slice []T) {

	h.data = append(h.data, slice...)
	for i := h.Size()/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *Heap[T]) Top() T {

	if h.Empty() {
		panic("cannot retrieve top value, heap is empty")
	}

	return h.data[0]
}

func (h *Heap[T]) Pop() {

	if h.Empty() {
		panic("cannot pop top value, heap is empty")
	}

	h.remove(0)
	h.heapify(0)
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) Empty() bool {
	return len(h.data) == 0
}

func (h Heap[T]) parent(index int) int {
	return (index - 1) / 2
}

func (h Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) remove(i int) {
	h.swap(i, h.Size()-1)
	h.data = h.data[:h.Size()-1]
}

func (h *Heap[T]) heapify(index int) {

	left := index*2 + 1
	right := index*2 + 2
	res := index

	if left < len(h.data) && h.cmp(h.data[left], h.data[index]) {
		res = left
	}

	if right < len(h.data) && h.cmp(h.data[right], h.data[res]) {
		res = right
	}

	if res != index {
		h.swap(res, index)
		h.heapify(res)
	}
}
