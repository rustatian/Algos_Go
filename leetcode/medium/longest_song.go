package medium

func numPairsDivisibleBy60(time []int) int {
	nums := [60]int{}
	res := 0

	for i := 0; i < len(time); i++ {
		if time[i]%60 == 0 {
			res += nums[0]
			continue
		}

		res += nums[60-time[i]%60]
		nums[time[i]%60]++
	}

	return res
}
