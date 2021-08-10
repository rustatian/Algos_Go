package medium

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {

	// Write your code here.
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
	// Write your code here.
}

func (h *MinHeap) siftUp() {
	// Write your code here.
}

func (h MinHeap) Peek() int {
	// Write your code here.
	return -1
}

func (h *MinHeap) Remove() int {
	// Write your code here.
	return -1
}

func (h *MinHeap) Insert(value int) {
	// Write your code here.
}
