package easy

func findDisappearedNumbers(nums []int) []int {
	missed := make([]int, 0, 2)

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n < 0 {
			n = n * -1
		}

		n = n - 1

		if nums[n] < 0 {
			continue
		}

		nums[n] = nums[n] * -1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}

		missed = append(missed, i+1)
	}

	return missed
}
