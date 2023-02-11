package medium

func wiggleSort(nums []int) {
	operation := 1
	for i := 0; i < len(nums)-1; i++ {
		switch operation {
		case 1:
			operation = 2
			if nums[i] <= nums[i+1] {
				continue
			} else {
				tmp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = tmp
			}
		case 2:
			operation = 1
			if nums[i] >= nums[i+1] {
				continue
			} else {
				tmp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = tmp
			}
		}
	}
}
