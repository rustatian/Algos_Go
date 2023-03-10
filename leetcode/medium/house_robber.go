package medium

func rob(nums []int) int {
	// up to k houses -> state variable
	// i -> state variable (loop)

	/*
		recurrence relation
		dp(i - 1) whatever money we can get if we wanted to rob curr

		rob -> nums[i] money (possible is we did not rob the prev house)

		(dp(i-1), dp(i-2) + nums[i])
	*/

	/*
		base cases:
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
	*/
	mem := make(map[int]int, 10)

	return dp(len(nums)-1, nums, &mem)
}

func dp(i int, nums []int, mem *map[int]int) int {
	if i == 0 {
		return nums[0]
	}

	if i == 1 {
		return robMax(nums[0], nums[1])
	}

	if val, ok := (*mem)[i]; ok {
		return val
	}

	(*mem)[i] = robMax(dp(i-1, nums, mem), dp(i-2, nums, mem)+nums[i])

	return (*mem)[i]
}

func robMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
