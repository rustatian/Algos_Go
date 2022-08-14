package easy

import (
	"math"
)

func sortedSquares(nums []int) []int {
	start := 0
	end := len(nums) - 1

	res := make([]int, len(nums))

	// we start iterating from the biggest positive (or negative smallest) number
	// and check 2 pointers
	// 1-st pointer (start) points to the smallest negative/positive number.
	// 2-nd pointer (end) points to the biggest positive/negative number.
	for i := len(nums) - 1; i >= 0; i-- {
		// if nums[start] by module is smaller that biggest positive number -> take a positive and decrement end ptr
		if math.Abs(float64(nums[start])) < math.Abs(float64(nums[end])) {
			res[i] = nums[end] * nums[end]
			end--
			continue
		}

		res[i] = nums[start] * nums[start]
		start++
	}

	return res
}
