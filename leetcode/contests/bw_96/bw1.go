package bw_96

func getCommon(nums1 []int, nums2 []int) int {
	idx1 := 0
	idx2 := 0

	for idx1 < len(nums1) && idx2 < len(nums2) {
		n1 := nums1[idx1]
		n2 := nums2[idx2]

		if n1 < n2 {
			idx1++
			continue
		}

		if n1 > n2 {
			idx2++
			continue
		}

		return nums1[idx1]
	}

	return -1
}
