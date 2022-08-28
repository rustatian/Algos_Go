package easy

func findMaxConsecutiveOnes(nums []int) int {
	num := 0
	curr := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if curr > num {
				num = curr
			}
			curr = 0
			continue
		}

		curr++
	}

	if curr > 0 {
		if curr > num {
			return curr
		}
	}
	return num
}
