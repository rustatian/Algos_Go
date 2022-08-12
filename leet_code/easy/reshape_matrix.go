package easy

func matrixReshape(mat [][]int, r int, c int) [][]int {
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	if len(mat) == 0 || r*c != len(mat)*len(mat[0]) {
		return mat
	}

	count := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			res[count/c][count%c] = mat[i][j]
			count++
		}

	}

	return res
}
