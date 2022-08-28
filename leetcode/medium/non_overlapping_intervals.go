package medium

import (
	"sort"
)

type vals2 [][]int

func (s vals2) Len() int {
	return len(s)
}

func (s vals2) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s vals2) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(vals2(intervals))

	end := intervals[0][1]
	count := 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}

	return len(intervals) - count
}
