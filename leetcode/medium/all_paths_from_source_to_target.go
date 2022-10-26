package medium

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph) - 1                 // target
	res := make([][]int, 0, len(graph)) // result
	// start from the 0
	currPath := []int{0}

	var dfsLocal func(cn int, path []int)

	dfsLocal = func(cn int, path []int) {
		if cn == n {
			cp := make([]int, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}

		for i := 0; i < len(graph[cn]); i++ {
			path = append(path, graph[cn][i])
			dfsLocal(graph[cn][i], path)
			path = path[:len(path)-1]
		}
	}

	dfsLocal(0, currPath)

	return res
}
