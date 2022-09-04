package easy

func removeDuplicates(nums []int) int {
	pWr := 0
	p1 := 1

	for p1 <= len(nums) - 1 {
		if nums[pWr] != nums[p1] {
			pWr++
			nums[pWr] = nums[p1]
		}
		p1++
	}

	return pWr + 1
}
