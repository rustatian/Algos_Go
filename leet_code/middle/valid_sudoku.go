package middle

func isValidSudoku(board [][]byte) bool {
	hm := make(map[byte]bool)
	hm2 := make(map[int]bool)
	// row
	for i := 0; i < len(board); i++ {
		// value in the row
		for j := 0; j < len(board[i]); j++ {
			// skip dots
			if board[i][j] == '.' {
				continue
			}

			if _, ok := hm[board[i][j]]; ok {
				return false
			} else {
				hm[board[i][j]] = true
			}
		}

		clear(hm)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// skip dots
			if board[j][i] == '.' {
				continue
			}

			if _, ok := hm[board[j][i]]; ok {
				return false
			} else {
				hm[board[j][i]] = true
			}
		}

		clear(hm)
	}

	// 0 0 0
	// 1 1 1
	// 2 2 2
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			//skip dots
			if board[r][c] == '.' {
				continue
			}

			br := (r/3)*3 + c/3
			bc := (r%3)*3 + c%3

			idx := int((board[br][bc] - '0') - 1)

			if _, ok := hm2[idx]; ok {
				return false
			}
			hm2[idx] = true

		}
	}

	return true
}

func clear2(hm2 map[int]byte) {
	for k := range hm2 {
		delete(hm2, k)
	}
}

func clear(hm map[byte]bool) {
	for k := range hm {
		delete(hm, k)
	}
}
