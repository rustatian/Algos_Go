package bw_86

import (
	"strconv"
)

func isStrictlyPalindromic(n int) bool {
	for i := 2; i < n-2+1; i++ {
		if !twoPointers(strconv.FormatInt(int64(i), 2)) {
			return false
		}
	}

	return true
}

func twoPointers(num string) bool {
	start := 0
	end := len(num) - 1

	for start < end {
		if num[start] == num[end] {
			start++
			end--
			continue
		}

		return false
	}

	return true
}
