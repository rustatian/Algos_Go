package bw_95

func xorBeauty(nums []int) int {
	beauty := 0

	for i := 0; i < len(nums); i++ {
		beauty ^= nums[i]
	}

	return beauty
}
