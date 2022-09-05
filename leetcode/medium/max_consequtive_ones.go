package medium

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	maxCurr := 0
	max := 0

	pSt := 0
	pRem := 0
	replaced := false

	for pSt <= len(nums)-1 {
		if nums[pSt] == 1 {
			pSt++
			maxCurr++
			continue
		}

		// replace done
		if !replaced {
			// remember position of the last 0
			pRem = pSt

			pSt++
			maxCurr++
			replaced = true
			continue
		}

		// start from the last 0
		pSt = pRem + 1
		if maxCurr > max {
			max = maxCurr
		}

		if len(nums[pSt:]) < max {
			return max
		}

		maxCurr = 0
		replaced = false
	}

	if maxCurr > max {
		return maxCurr
	}

	return max
}
