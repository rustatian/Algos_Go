package medium

func permute(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	n := len(nums)

	backtrack(0, n, nums, &res)

	return res
}

func backtrack(first, n int, nums []int, res *[][]int) {
	if first == n {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp[:])
	}

	for i := first; i < n; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtrack(first+1, n, nums, res)
		nums[first], nums[i] = nums[i], nums[first]
	}
}
