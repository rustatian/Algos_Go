package bw_97

func maxCount(banned []int, n int, maxSum int) int {
	ban := make(map[int]struct{}, 10)
	count := 0

	for i := 0; i < len(banned); i++ {
		ban[banned[i]] = struct{}{}
	}

	sum := 0

	for i := 1; i <= n; i++ {
		if _, ok := ban[i]; ok {
			continue
		}

		sum += i

		if sum > maxSum {
			return count
		}

		count++
	}

	return count
}
