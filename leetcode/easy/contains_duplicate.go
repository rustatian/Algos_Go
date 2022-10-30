package easy

func containsNearbyDuplicate(nums []int, k int) bool {
	hm := make(map[int]struct{}, 5)

	for i := 0; i < len(nums); i++ {
		if _, ok := hm[nums[i]]; ok {
			return true
		}

		hm[nums[i]] = struct{}{}
		if len(hm) > k {
			delete(hm, nums[i-k])
		}
	}

	return false
}
