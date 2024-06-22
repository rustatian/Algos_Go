package bw_133

func minOperations2(nums []int) int {
	n := len(nums)
	flips := 0
	flipped := false

	for i := 0; i < n; i++ {
		current := nums[i]
		if flipped {
			current ^= 1
		}
		if current == 0 {
			flips++
			flipped = !flipped
		}
	}

	return flips
}
