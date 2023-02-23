package easy

import (
	"fmt"
)

func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		return nil
	}

	res := make([]string, 0, 10)

	prev := lower - 1
	curr := -1
	for i := 0; i < len(nums)+1; i++ {
		if i < len(nums) {
			curr = nums[i]
		} else {
			curr = upper + 1
		}

		if prev+1 <= curr-1 {
			if prev+1 == curr-1 {
				res = append(res, fmt.Sprintf("%d", prev+1))
			} else {
				// prev+1 -> low
				// curr-1 -> curr
				res = append(res, fmt.Sprintf("%d->%d", prev+1, curr-1))
			}
		}

		prev = curr
	}

	return res
}
