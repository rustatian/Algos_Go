package easy

func pivotIndex(nums []int) int {
	var sum, left = 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if left == sum-left-nums[i] {
			return i
		}

		left += nums[i]
	}

	return -1
}
