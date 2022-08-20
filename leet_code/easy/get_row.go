package easy

func getRow(rowIndex int) []int {
	res := make([][]int, 33)

	// C(n,k) = C(n-1,k-1) + C(n-1,k)
	for i := 0; i < rowIndex+1; i++ {
		for k := 0; k < i+1; k++ {
			res[i] = append(res[i], 1)
		}
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res[rowIndex]
}
