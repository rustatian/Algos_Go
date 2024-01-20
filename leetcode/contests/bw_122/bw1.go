package bw_122

import (
	"sort"
)

func minimumCost(nums []int) int {
	if len(nums) == 0 || len(nums) < 3 {
		return 0
	}

	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums[1:])

	return nums[0] + nums[1] + nums[2]
}
