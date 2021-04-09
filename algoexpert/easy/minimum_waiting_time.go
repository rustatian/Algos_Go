package main

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	sum := 0
	res := 0

	for i := 1; i < len(queries); i++ {
		sum += queries[i-1]
		res += sum
	}

	// Write your code here.
	return res
}
