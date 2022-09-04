package easy

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	pWr := 0
	pEl := 1

	for pEl <= len(nums)-1 {
		if nums[pWr] == 0 {
			if nums[pEl] != 0 {
				nums[pWr] = nums[pEl]
				nums[pEl] = 0
				pWr++
			}
		} else {
			pWr++
		}
		pEl++
	}
}
