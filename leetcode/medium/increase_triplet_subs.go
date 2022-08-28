package medium

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	num1 := 1<<((32<<(^uint(0)>>63))-1) - 1
	num2 := num1

	for i := 0; i < len(nums); i++ {
		if nums[i] <= num1 {
			num1 = nums[i]
		} else if nums[i] <= num2 {
			num2 = nums[i]
		} else {
			return true
		}
	}

	return false
}
