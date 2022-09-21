package medium

import (
	"sort"
)

func earliestAcq(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	root := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = 1
	}

	steps := n

	for i := 0; i < len(logs); i++ {
		steps -= union4(logs[i][1], logs[i][2], root, rank)

		if steps == 1 {
			return logs[i][0]
		}
	}

	return -1
}

func find4(x int, root []int) int {
	if root[x] == x {
		return x
	}

	root[x] = find4(root[x], root)

	return root[x]
}

func union4(x, y int, root []int, rank []int) int {
	rootX := find4(x, root)
	rootY := find4(y, root)

	// already friends
	if rootX == rootY {
		return 0
	}

	if rank[rootX] > rank[rootY] {
		root[rootY] = rootX
	} else if rank[rootX] < rank[rootY] {
		root[rootX] = rootY
	} else {
		root[rootY] = rootX
		rank[rootX]++
	}

	// add friend
	return 1
}
