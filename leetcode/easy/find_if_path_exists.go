package easy

func validPath(n int, edges [][]int, source int, destination int) bool {
	root := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 0
	}

	for i := 0; i < len(edges); i++ {
		union(edges[i][0], edges[i][1], root, rank)
	}

	return find(root[source], root) == find(root[destination], root)
}

func find(x int, root []int) int {
	if x == root[x] {
		return x
	}

	root[x] = find(root[x], root)

	return root[x]
}

func union(x, y int, root []int, rank []int) int {
	rootX := find(x, root)
	rootY := find(y, root)

	if rootX == rootY {
		// union
		return rootX
	}

	if rank[rootX] > rank[rootY] {
		t := rootX
		rootX = rootY
		rootY = t
	}

	// connect
	root[rootX] = rootY
	rank[rootY] += rank[rootX]

	return rootY
}
