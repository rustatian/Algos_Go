package medium

func sortColors(nums []int) {
	p0 := 0
	p2 := len(nums) - 1

	curr := 0

	for curr <= p2 {
		switch nums[curr] {
		case 0:
			tmp := nums[curr]
			nums[curr] = nums[p0]
			nums[p0] = tmp
			p0++
			curr++
		case 1:
			curr++
		case 2:
			tmp := nums[p2]
			nums[p2] = nums[curr]
			nums[curr] = tmp
			p2--
		default:
			panic("should be only 0, 1 and 2")
		}
	}
}
