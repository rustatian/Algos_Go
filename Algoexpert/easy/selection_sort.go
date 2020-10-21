package main

import (
	"math"
)

func SelectionSort(array []int) []int {
	smallest := math.MaxInt64
	index := 0
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				if array[j] < smallest {
					smallest = array[j]
					index = j
				}
			}
		}
		// we actually at the smallest number
		if smallest == math.MaxInt64 {
			continue
		}
		tmp := array[i]
		array[i] = smallest
		array[index] = tmp
		smallest = math.MaxInt64
	}

	return array
}
