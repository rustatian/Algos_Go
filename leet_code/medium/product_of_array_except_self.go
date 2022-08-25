package medium

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	L := make([]int, len(nums))
	L[0] = 1

	for i := 0; i < len(res); i++ {
		res[i] = nums[i] * nums[i]
	}

	post := 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= post
		post *= res[i]
	}

	return res
}
