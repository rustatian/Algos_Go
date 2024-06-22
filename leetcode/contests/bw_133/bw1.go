package bw_133

func minimumOperations(nums []int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		rem := nums[i] % 3
		if rem == 0 {
			continue
		}

		if rem == 1 || rem == 2 {
			res += 1
		}
	}

	return res
}
