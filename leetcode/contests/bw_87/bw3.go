package bw_87

func smallestSubarrays(nums []int) []int {
	res := make([]int, len(nums))

	OR := 0

	for i := 0; i < len(nums); i++ {
		OR |= nums[i]
	}

	for i := 0; i < len(nums); i++ {
		or := 0
		for j := i; j < len(nums); j++ {
			or |= nums[j]
			if or >= OR {
				res[i] = len(nums[i:j]) + 1
				break
			}

			res[i] = len(nums[i:j]) + 1
		}
	}

	res[len(res)-1] = 1
	return res
}
