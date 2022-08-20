package medium

func rotate(matrix [][]int) {
	rows := len(matrix)

	for i := 0; i < rows/2; i++ {
		for j := i; j < rows-i-1; j++ {
			tmp := matrix[rows-1-j][i]

			matrix[rows-1-j][i] = matrix[rows-1-i][rows-j-1]
			matrix[rows-1-i][rows-j-1] = matrix[j][rows-1-i]
			matrix[j][rows-1-i] = matrix[i][j]
			matrix[i][j] = tmp
		}
	}
}
