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

// NewHeap returns a new heap
// A heap can be built from an array in 3 methods:
// - Sort the array i.e. sort.Slice(arr, isSmallerFunc) ==> RUNTIME Θ(n*log(n))
// - for (i:=2; i <= size; i ++) heapifyUp(i)           ==> RUNTIME Θ(n*log(n))
// - for (i:=size/2; i>0; i--) heapifyDown(i)           ==> RUNTIME Θ(n)
// We will be using the third and most efficient way. Why is it Θ(n)?
// heapifyDown can only take one path down and will be ran n/2 times
// whereas heapifyUp will happen for every leaf and thus is ran n times
func NewHeap(arr []interface{}, isSmallerFunc func(int, int) bool) *AHeap {
	h := &AHeap{
		arr:       arr,
		isSmaller: isSmallerFunc,
		size:      len(arr),
	}
	// leaves always respect the max or min heap property
	// and so we only have to heapify down all nodes which are on the second
	// level from the bottom
	for i := h.size / 2; i > 0; i-- {
		h.heapifyDown(i)
	}
	return h
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

// IntsIsSmallerFunc is an example isSmallerFunc for integers
func IntsIsSmallerFunc(a, b int) bool {
	return a < b
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
		// i is now at the index which used to be the smallest's
		h.heapifyDown(smallest)
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
