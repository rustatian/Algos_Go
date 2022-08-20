package medium

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	res := make([][]int, m)
	res[0] = make([]int, n+1)
	for i := 0; i < len(res[0]); i++ {
		res[0][i] = 1
	}

	// cols
	for i := 1; i < m; i++ {
		res[i] = make([]int, n+1)
		// rows
		for j := 1; j < n+1; j++ {
			res[i][j] = res[i][j-1] + res[i-1][j]
		}
	}

	return res[m-1][n]
}
