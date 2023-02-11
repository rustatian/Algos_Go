package medium

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	distance := 0

	// first jump
	startJump := nums[0]
	// first jump res
	res := 1

	for i := 1; i < len(nums)-1; i++ {
		// max jump from the next
		distance = maxx(distance, i+nums[i])
		// if we can jump further, we should start on a next iteration from the far-far-far index
		if i == startJump {
			startJump = distance
			res++
		}
	}

	return res
}

func maxx(a, b int) int {
	if a > b {
		return a
	}

	return b
}
