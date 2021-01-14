package main

func IsValidSubsequence(array []int, sequence []int) bool {
	if len(array) == 1 && len(sequence) == 1 && array[0] == sequence[0] {
		return true
	}

	idx := 0
	for i := 0; i < len(array); i++ {
		if array[i] == sequence[idx] {
			idx++
			if idx == len(sequence) {
				break
			}
		}
	}

	if idx == len(sequence) {
		return true
	}
	return false

}
