package easy

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}

		return 1
	}

	pWr := 0
	pl := len(nums) - 1

	for pWr <= pl {
		if nums[pWr] == val {
			if nums[pl] != val {
				nums[pWr] = nums[pl]
				pWr++
				pl--
			} else {
				pl--
			}
		} else {
			pWr++
		}
	}

	if pWr <= 0 {
		return 0
	}

	return pWr
}
