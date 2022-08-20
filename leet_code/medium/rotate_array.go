package medium

// > 20m, see

func rotate2(nums []int, k int) {
	k %= len(nums)
	// reverse the whole slice
	rev(nums, 0, len(nums)-1)
	// reverse the elements up to k
	rev(nums, 0, k-1)
	// reverse back the rest of the slice
	rev(nums, k, len(nums)-1)
}

func rev(nums []int, start, end int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}
