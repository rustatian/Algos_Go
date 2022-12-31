package bw_87

import (
	"sort"
)

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	hm := make(map[float64]struct{})

	st := 0
	end := len(nums) - 1

	for st <= end {
		n1 := float64(nums[st])
		n2 := float64(nums[end])
		hm[(n1+n2)/2] = struct{}{}
		st++
		end--
	}

	return len(hm)
}
