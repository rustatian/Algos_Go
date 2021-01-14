package medium

func SpiralTraverse(array [][]int) []int {
	res := make([]int, 0, 0)
	startRow := 0
	endRow := len(array) - 1
	startColumn := 0
	endColumn := len(array[0]) - 1

	for startRow <= endRow && startColumn <= endColumn {
		for i := startColumn; i <= endColumn; i++ {
			res = append(res, array[startRow][i])
		}
		for i := startRow + 1; i <= endRow; i++ {
			res = append(res, array[i][endColumn])
		}

		for i := endColumn - 1; i >= startColumn; i-- {
			if startRow == endRow {
				break
			}
			res = append(res, array[endRow][i])
		}
		for i := endRow - 1; i > startRow; i-- {
			if startColumn == endColumn {
				break
			}
			res = append(res, array[i][startColumn])
		}

		startColumn += 1
		endColumn -= 1
		startRow += 1
		endRow -= 1
	}

	return res
}
