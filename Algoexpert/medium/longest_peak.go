package medium

func LongestPeak(array []int) int {
	for i := 1; i < len(array); i++ {
		peak := array[i-1] < array[i] && array[i] > array[i+1]
		if !peak {
			continue
		}

		left := i

	}
	return -1
}
