package easy

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	cIdx := len(nums) - 1 // initial c
	aIdx := len(nums) - 2 // initial b
	bIdx := len(nums) - 3 // initial a

	for bIdx >= 0 {
		a := nums[aIdx]
		b := nums[bIdx]
		c := nums[cIdx]

		if c < a+b {
			return a + b + c
		}

		cIdx--
		aIdx--
		bIdx--
	}

	return 0
}
