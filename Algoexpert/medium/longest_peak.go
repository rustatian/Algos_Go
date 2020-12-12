package medium

func LongestPeak(array []int) int {
	curPeak := 0
	for i := 1; i < len(array)-1; i++ {
		peak := array[i-1] < array[i] && array[i] > array[i+1]
		if !peak {
			continue
		}

		left := i - 2
		for j := left; j >= 0 && array[j] < array[j+1]; {
			j--
			left = j
		}

		right := i + 2
		for k := right; k < len(array) && array[k] < array[k-1]; {
			k++
			right = k
		}

		length := right - left - 1
		if length > curPeak {
			curPeak = length
		}
		i = right - 1
	}
	return curPeak
}
