package easy

// 4m42s
// rust -> hashmap

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0, 1)

	var p1 = 0
	var p2 = 0

	for {
		if p1 > len(nums1)-1 {
			return res
		}

		if p2 > len(nums2)-1 {
			return res
		}

		p1el := nums1[p1]
		p2el := nums2[p2]

		if p1el == p2el {
			res = append(res, p1el)
			p1++
			p2++
			continue
		}

		if p1el < p2el {
			p1++
			continue
		}

		if p1el > p2el {
			p2++
		}
	}
}
