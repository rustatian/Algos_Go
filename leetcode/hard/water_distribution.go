package hard

import (
	"sort"
)

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	group := make([]int, n+1)
	rank := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		group[i] = i
		rank[i] = 0
	}

	edges := make([][]int, n+1+len(pipes))
	for i:=0;i<len(edges);i++ {
		edges[i] = make([]int, 3)
	}

	for i := 0; i < len(wells); i++ {
		edges = append(edges, []int{0, i + 1, wells[i]})
	}

	for i := 0; i < len(pipes); i++ {
		edges = append(edges, pipes[i])
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	tcost := 0

	for i := 0; i < len(edges); i++ {
		h1 := edges[i][0]
		h2 := edges[i][1]
		cost := edges[i][2]

		if union1(h1, h2, group, rank) {
			tcost += cost
		}
	}

	return tcost
}

func find1(p int, group []int) int {
	if group[p] != p {
		group[p] = find1(group[p], group)
	}

	return group[p]
}

func union1(p1, p2 int, groups []int, ranks []int) bool {
	g1 := find1(p1, groups)
	g2 := find1(p2, groups)

	if g1 == g2 {
		return false
	}

	if ranks[g1] > ranks[g2] {
		groups[g2] = g1
	} else if ranks[g1] < ranks[g2] {
		groups[g1] = g2
	} else {
		groups[g1] = g2
		ranks[g2] += 1
	}

	return true
}
