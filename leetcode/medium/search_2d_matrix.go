package medium

func searchMatrix(matrix [][]int, target int) bool {

	row := len(matrix) - 1
	col := 0

	for col <= len(matrix[0])-1 && row >= 0 {
		switch {
		case matrix[row][col] > target:
			row--
		case matrix[row][col] < target:
			col++
		default:
			return true
		}
	}

	return false
}
