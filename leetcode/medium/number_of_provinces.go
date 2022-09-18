package medium

func findCircleNum(isConnected [][]int) int {
	N := len(isConnected[0])

	root := make([]int, N)

	for i := 0; i < N; i++ {
		root[i] = -1
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			// connected
			if isConnected[i][j] == 1 {
				if i != j {
					// i-th city to the j-th
					union1(i, j, &root)
				}
			}
		}
	}

	count := 0
	for i := 0; i < N; i++ {
		// not connected
		if root[i] == -1 {
			count++
		}
	}

	// if all connected to each other, the last city connection will be -1
	return count
}

func find1(x int, root []int) int {
	if root[x] == -1 {
		// connected to self
		return x
	}

	return find1(root[x], root)
}

func union1(x, y int, root *[]int) {
	rootX := find1(x, *root)
	rootY := find1(y, *root)

	if rootX != rootY {
		(*root)[rootX] = rootY
	}
}
