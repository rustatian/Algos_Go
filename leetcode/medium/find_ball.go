package medium

func findBall(grid [][]int) []int {
	// cols
	res := make([]int, 0, len(grid[0]))

	for i := 0; i < len(grid[0]); i++ {
		res = append(res, check(0, i, grid))
	}

	return res
}

func check(row, col int, grid [][]int) int {
	if row == len(grid) {
		return col
	}

	nc := col + grid[row][col]

	// out of boundaries
	if nc < 0 || nc > len(grid[0]) {
		return -1
	}

	if grid[row][col] != grid[row][nc] {
		return -1
	}

	return check(row+1, nc, grid)
}
