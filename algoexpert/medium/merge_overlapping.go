package medium

import "sort"

type Sorter struct {
	intervals [][]int
}

func (t Sorter) Len() int {
	return len(t.intervals)
}

func (t Sorter) Less(i, j int) bool {
	return t.intervals[i][0] < t.intervals[j][0]
}

func (t Sorter) Swap(i, j int) {
	t.intervals[i], t.intervals[j] = t.intervals[j], t.intervals[i]
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	a := Sorter{intervals: intervals}
	sort.Sort(a)

loop:
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][1] < intervals[i+1][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			goto loop
		}
	}
	return intervals
}
