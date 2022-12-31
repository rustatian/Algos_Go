package hard

import (
	"container/heap"
)

type MedianFinder struct {
	s *minHeap
	l *maxHeap
}

func Constructor() MedianFinder {
	s := &minHeap{}
	l := &maxHeap{}

	heap.Init(s)
	heap.Init(l)
	return MedianFinder{
		s: s,
		l: l,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.s, num)
	heap.Push(mf.l, (*mf.s)[0])
	heap.Pop(mf.s)

	if mf.s.Len() < mf.l.Len() {
		heap.Push(mf.s, (*mf.l)[0])
		heap.Pop(mf.l)
	}

}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.s.Len() > mf.l.Len() {
		return float64((*mf.s)[0])
	}

	return (float64((*mf.s)[0]) + float64((*mf.l)[0])) / 2.0
}

type minHeap []int

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Pop() any {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func (m *minHeap) Push(item any) {
	*m = append(*m, item.(int))
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

type maxHeap []int

func (m *maxHeap) Len() int {
	return len(*m)
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Pop() any {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func (m *maxHeap) Push(item any) {
	*m = append(*m, item.(int))
}
