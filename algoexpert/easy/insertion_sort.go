package main

func InsertionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
