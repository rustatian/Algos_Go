package medium

func SpiralTraverse(array [][]int) []int {
	res := make([]int, 0, 0)
	startRow := 0
	startColumn := 0
	endRow := len(array)
	endColumn := len(array[0]) - 1

	for endRow != 0 && endColumn != 0 {
		res = append(res, array[startRow][startColumn])
		startColumn++
		if startColumn == endColumn {
			for startRow != endRow {
				res = append(res, array[startRow][startColumn])
				startRow++
			}

		}
	}
	return res
}
