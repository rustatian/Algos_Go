package medium

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	res := make([][]int, 0, len(edges)) // result
	currPath := []int{source}

	var dfsLocal func(cn int, path []int)

	dfsLocal = func(cn int, path []int) {
		if cn == destination {
			cp := make([]int, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}

		for i := 0; i < len(edges[cn]); i++ {
			path = append(path, edges[cn][i])
			dfsLocal(source, path)
			path = path[:len(path)-1]
		}
	}

	dfsLocal(0, currPath)

	return len(res) > 0
}
