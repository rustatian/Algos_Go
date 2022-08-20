package medium

import (
	"sort"
)

type vals [][]int

func (s vals) Len() int {
	return len(s)
}

func (s vals) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s vals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(vals(intervals))

	res := make([][]int, 0, len(intervals))

	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
