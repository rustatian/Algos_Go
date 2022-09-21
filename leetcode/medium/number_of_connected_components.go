package medium

func countComponents(n int, edges [][]int) int {
	root := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 1
	}

	count := n
	for i := 0; i < len(edges); i++ {
		count -= union3(edges[i][0], edges[i][1], root, rank)
	}

	return count
}

func find3(x int, root []int) int {
	if x == root[x] {
		return x
	}

	root[x] = find3(root[x], root)

	return root[x]
}

func union3(x, y int, root []int, rank []int) int {
	rootX := find3(x, root)
	rootY := find3(y, root)

	if rootX == rootY {
		return 0
	}

	if rank[rootX] > rank[rootY] {
		root[rootY] = rootX
	} else if rank[rootX] < rank[rootY] {
		root[rootX] = rootY
	} else {
		root[rootY] = rootX
		rootX++
	}

	return 1
}
