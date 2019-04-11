package heap

// Heap is a functional representation of a heap
type Heap interface {
	Insert(interface{})
	RemoveMin() interface{}
}

// AHeap is a heap implemented with an array
type AHeap struct {
	arr       []interface{}
	isSmaller func(int, int) bool
	size      int
}

// Insert inserts an item into the heap
func (h *AHeap) Insert(x interface{}) {
	if h.size == len(h.arr) {
		h.arr = append(h.arr, make([]interface{}, len(h.arr))...)
	}
	h.size++
	h.arr[h.size] = x
	h.heapifyUp(h.size)
}

// RemoveMin removes the minimum element from a heap
func (h *AHeap) RemoveMin() interface{} {
	if h.isEmpty() {
		return nil
	}
	// smallest is the root of the heap
	min := h.arr[1]
	h.size--
	h.heapifyDown(1)
	return min
}

// heapifyDown moves the node at index i down to its appropriate spot
func (h *AHeap) heapifyDown(i int) {
	smallest := i
	left := i * 2
	right := left + 1
	if (left <= h.size) && h.isSmaller(left, smallest) {
		smallest = left
	}
	if (left <= h.size) && h.isSmaller(right, smallest) {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.heapifyUp(smallest)
	}
}

// heapifyDown moves the node at index i up to its appropriate spot
func (h *AHeap) heapifyUp(i int) {
	if i == 1 {
		return
	}
	parent := i / 2
	if h.isSmaller(i, parent) {
		h.swap(i, parent)
		// i is now at the index which used to be the parent's
		h.heapifyUp(parent)
	}
}

func (h *AHeap) swap(a, b int) {
	tmp := h.arr[a]
	h.arr[a] = h.arr[b]
	h.arr[b] = tmp
}

func (h *AHeap) isEmpty() bool {
	return (h.size == 1)
}
