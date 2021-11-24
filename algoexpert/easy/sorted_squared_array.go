package main

import (
	"sort"
)

func SortedSquaredArray(array []int) []int {
	res := make([]int, len(array))

	for i := 0; i < len(array); i++ {
		res[i] = array[i] * array[i]
	}

	sort.Ints(res)

	return res
}
