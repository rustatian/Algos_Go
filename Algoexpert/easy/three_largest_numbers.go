package main

import "math"

func FindThreeLargestNumbers(array []int) []int {
	res := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for i := 0; i < len(array); i++ {
		if array[i] > res[2] {
			res[0] = res[1]
			res[1] = res[2]
			res[2] = array[i]
		} else if array[i] > res[1] {
			res[0] = res[1]
			res[1] = array[i]
		} else if array[i] > res[0] {
			res[0] = array[i]
		}
	}
	return res
}
