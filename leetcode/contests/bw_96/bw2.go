package bw_96

import (
	"math"
)

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	if len(nums1) == 0 {
		return -1
	}

	noDiff := true
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			noDiff = false
			break
		}
	}

	if noDiff {
		return 0
	}

	pos := 0
	neg := 0

	for i := 0; i < len(nums1); i++ {
		if math.Mod(math.Abs(float64(nums2[i]-nums1[i])), float64(k)) != 0 {
			return -1
		}

		val := nums2[i] - nums1[i]
		if val == 0 {
			continue
		}
		if val > 0 {
			pos += val
			continue
		}
		neg += val
	}

	if neg+pos != 0 {
		return -1
	}

	return int64(pos / k)
}
