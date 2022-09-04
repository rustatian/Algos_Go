package main

import (
	"sort"
	"testing"
)

func BenchmarkHeapSort(b *testing.B) {
	a := []int{2, 23, 33, 44, 1, 2, 2, 2, 4, 6, 99}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		heapSort(a)
	}
}

func BenchmarkHeapSortStd(b *testing.B) {
	a := []int{2, 23, 33, 44, 1, 2, 2, 2, 4, 6, 99}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})
	}
}
