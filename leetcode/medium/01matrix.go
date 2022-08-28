package medium

import (
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	if rows == 0 {
		return nil
	}

	cols := len(mat[0])

	vec := make([][]int, rows)

	for i := 0; i < rows; i++ {
		vec[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			vec[i][j] = math.MaxInt
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				vec[i][j] = 0
			} else {
				if i > 0 {
					vec[i][j] = min(vec[i][j], vec[i-1][j]+1)
				}
				if j > 0 {
					vec[i][j] = min(vec[i][j], vec[i][j-1]+1)
				}
			}
		}
	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i < rows-1 {
				vec[i][j] = min(vec[i][j], vec[i+1][j]+1)
			}

			if j < cols-1 {
				vec[i][j] = min(vec[i][j], vec[i][j+1]+1)
			}
		}
	}

	return vec
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
