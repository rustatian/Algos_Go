package medium

import (
	"math"
	"sort"
)

// optimal O((nlog(n) + mlog(m)) + O(1) time
// https://en.wikipedia.org/wiki/Absolute_difference
func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)

	left := 0
	right := 0

	smallest := math.MaxInt64
	res := make([]int, 2, 2)

	for left != len(array1) && right != len(array2) {
		diff := (array1[left]) - (array2[right])
		if diff < 0 {
			diff = diff * -1
		}
		if (diff) < smallest {
			res[0] = array1[left]
			res[1] = array2[right]
			smallest = diff
		}

		if array1[left]-array2[right] == 0 {
			res[0] = array1[left]
			res[1] = array2[right]
			return res
		} else if array1[left] < array2[right] {
			left++
			continue
		} else {
			right++
			continue
		}

	}

	return res
}
