package easy

func sortArrayByParity(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	pCurr := 0
	pCurrPl1 := 1

	for pCurrPl1 <= len(nums)-1 {
		if nums[pCurr]%2 != 0 {
			if nums[pCurrPl1]%2 == 0 {
				nums[pCurr], nums[pCurrPl1] = nums[pCurrPl1], nums[pCurr]
				pCurr++
			}
		} else {
			pCurr++
		}
		pCurrPl1++
	}

	return nums
}
