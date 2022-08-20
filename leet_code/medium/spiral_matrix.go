package medium

func generateMatrix(n int) [][]int {
	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// total number of layers
	layers := (n + 1) / 2
	num := 1

	for i := 0; i < layers; i++ {
		for j := i; j < n-i; j++ {
			res[i][j] = num
			num++
		}

		for j := i + 1; j < n-i; j++ {
			res[j][n-i-1] = num
			num++
		}

		for j := n - i - 2; j >= i; j-- {
			res[n-i-1][j] = num
			num++
		}

		for j := n - i - 2; j > i; j-- {
			res[j][i] = num
			num++
		}

	}

	return res
}
