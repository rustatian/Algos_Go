package medium

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) == 1 {
		return false
	}

	hm := make(map[int]int, len(nums))
	hm[0] = 0 // initial sum
	// prev val
	prev := 0
	for i := 0; i < len(nums); i++ {
		prev = prev + nums[i]
		if _, ok := hm[prev%k]; !ok {
			// just count
			hm[prev%k] = i + 1
			continue
		}

		if hm[prev%k] < i {
			return true
		}
	}

	return false
}
