package medium

type rotten struct {
	row int
	col int
}

func orangesRotting(grid [][]int) int {
	fresh := 0
	rows := len(grid)
	cols := len(grid[0])

	queue := make([]*rotten, 0, 1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, &rotten{
					row: i,
					col: j,
				})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	queue = append(queue, &rotten{
		row: -1,
		col: -1,
	})

	minutes := -1
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.row == -1 {
			minutes++
			if len(queue) > 0 {
				queue = append(queue, &rotten{
					row: -1,
					col: -1,
				})
			}
		} else {
			for i := 0; i < len(directions); i++ {
				nr := p.row + directions[i][0]
				nc := p.col + directions[i][1]

				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					if grid[nr][nc] == 1 {
						grid[nr][nc] = 2
						fresh--
						queue = append(queue, &rotten{
							row: nr,
							col: nc,
						})
					}
				}
			}
		}
	}

	if fresh == 0 {
		return minutes
	}

	return -1
}
