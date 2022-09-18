package medium

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	rank := make([]int, n)
	root := make([]int, n)

	for i := 0; i < n; i++ {
		root[i] = i
		// initial rank is 1 (standalone)
		rank[i] = 1
	}

	for ai := 0; ai < len(edges); ai++ {
		if !union2(edges[ai][0], edges[ai][1], &root, &rank) {
			return false
		}
	}

	return true
}

func find2(x int, root []int) int {
	if x == root[x] {
		return x
	}

	root[x] = find2(root[x], root)

	return root[x]
}

func union2(x, y int, root *[]int, rank *[]int) bool {
	rootX := find2(x, *root)
	rootY := find2(y, *root)

	if rootX == rootY {
		return false
	}

	if (*rank)[rootX] > (*rank)[rootY] {
		(*root)[rootY] = rootX
	} else if (*rank)[rootX] < (*rank)[rootY] {
		(*root)[rootX] = rootY
	} else {
		(*root)[rootY] = rootX
		(*rank)[rootX] += 1
	}

	return true
}
