package easy

import (
	"sort"
)

func findErrorNums(nums []int) []int {
	if len(nums) == 2 {
		nums[0] = nums[1] - 1
		return nums
	}

	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {

		}
	}

	return nums
}
