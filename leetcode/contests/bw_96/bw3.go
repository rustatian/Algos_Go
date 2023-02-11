package bw_96

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func (m *minHeap) Push(item interface{}) {
	*m = append(*m, item.(int))
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	nn := make([][2]int, n)

	for i := 0; i < n; i++ {
		nn[i] = [2]int{nums2[i], nums1[i]}
	}

	sort.Slice(nn, func(i, j int) bool {
		return nn[i][0] > nn[j][0]
	})

	mh := &minHeap{}
	heap.Init(mh)
	res := 0
	sum := 0

	for i := 0; i < n; i++ {
		heap.Push(mh, nn[i][1])
		sum += nn[i][1]

		if mh.Len() > k {
			el := heap.Pop(mh).(int)
			sum -= el
		}

		if mh.Len() == k {
			res = max(res, sum*nn[i][0])
		}
	}

	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
