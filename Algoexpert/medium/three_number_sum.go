package medium

import (
	"sort"
)

// pointers move approach
func ThreeNumberSum(array []int, target int) [][]int {
	// numbers distinct, sort
	sort.Ints(array)
	// result
	res := make([][]int, 0, 2)
	for i := 0; i < len(array); i++ {
		left := i + 1
		right := len(array) - 1
		for ; left < right; {
			current := array[i] + array[left] + array[right]
			if current == target {
				res = append(res, []int{array[i], array[left], array[right]})
				left++
				right--
			} else if current < target {
				left++
			} else if current > target {
				right--
			}
		}
	}
	if len(res) < 3 {
		return res
	}
	return res
}
