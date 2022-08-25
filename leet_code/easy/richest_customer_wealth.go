package easy

func maximumWealth(accounts [][]int) int {
	res := -1

	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}

		if sum > res {
			res = sum
		}
	}

	return res
}
