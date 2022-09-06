package easy

import (
	"math"
)

func thirdMax(nums []int) int {
	res := make(map[int]struct{}, 4)

	for i := 0; i < len(nums); i++ {
		res[nums[i]] = struct{}{}
		if len(res) > 3 {
			remove(&res)
		}
	}

	max := math.MinInt
	if len(res) < 3 {
		for k := range res {
			if k > max {
				max = k
			}
		}

		return max
	}

	max = math.MaxInt
	for k := range res {
		if k < max {
			max = k
		}
	}

	return max
}

func remove(nums *map[int]struct{}) {
	val := math.MaxInt

	for k := range *nums {
		if k < val {
			val = k
		}
	}

	delete(*nums, val)
}
