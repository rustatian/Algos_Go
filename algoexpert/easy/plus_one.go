package main

func plusOne(digits []int) []int {
	if len(digits) == 1 {
		if digits[0] == 9 {
			return []int{1, 0}
		}

		digits[0]++
		return digits
	}

	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}

	hasRem := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			hasRem = true
			continue
		} else {
			digits[i]++
			hasRem = false
			return digits
		}
	}

	if hasRem {
		res := make([]int, len(digits)+1)
		res[0] = 1
		return res
	}

	return digits
}
