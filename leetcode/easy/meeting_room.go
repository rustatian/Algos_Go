package easy

import (
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Sort(inter(intervals))

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}

	return true
}

type inter [][]int

func (s inter) Len() int {
	return len(s)
}

func (s inter) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s inter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
