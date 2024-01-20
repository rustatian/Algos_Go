package bw_122

import (
	"sort"
)

func canSortArray(nums []int) bool {
	countBits := func(x int) int {
		count := 0
		for x > 0 {
			count += x & 1
			x >>= 1
		}
		return count
	}

	groups := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		nb1 := countBits(nums[i])
		groups[nb1] = append(groups[nb1], nums[i])
	}

	for _, v := range groups {
		sort.Ints(v)

		for i := 0; i < len(v)-1; i++ {
			if v[i] >= v[i+1] {
				return false
			}
		}
	}

	return true
}
