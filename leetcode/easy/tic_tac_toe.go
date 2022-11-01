package easy

import (
	"math"
)

func tictactoe(moves [][]int) string {
	bs := 3 // board size
	rows := make([]int, bs)
	cols := make([]int, bs)

	var diag, otherDiag = 0, 0

	// 1 - A
	// 2 - B

	var player = 1 // A is first

	for i := 0; i < len(moves); i++ {
		row := moves[i][0]
		col := moves[i][1]

		rows[row] += player
		cols[col] += player

		if row == col {
			diag += player
		}

		if row+col == bs-1 {
			otherDiag += player
		}

		if math.Abs(float64(rows[row])) == float64(bs) || math.Abs(float64(cols[col])) == float64(bs) || math.Abs(float64(diag)) == float64(bs) || math.Abs(float64(otherDiag)) == float64(bs) {
			if player == 1 {
				return "A"
			}

			return "B"
		}

		player *= -1
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}
