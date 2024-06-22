package bw_133

func minOperations(nums []int) int {
	n := len(nums)
	flips := 0

	for i := 0; i <= n-3; i++ {
		if nums[i] == 0 {
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			flips++
		}
	}

	for i := n - 2; i < n; i++ {
		if nums[i] == 0 {
			return -1
		}
	}

	return flips
}
