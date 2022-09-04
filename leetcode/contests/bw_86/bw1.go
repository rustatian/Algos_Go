package bw_86

func findSubarrays(nums []int) bool {
	hm := make(map[int]struct{})

	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i] + nums[i+1]
		if _, ok := hm[sum]; ok {
			return true
		} else {
			hm[sum] = struct{}{}
		}
	}

	return false
}
