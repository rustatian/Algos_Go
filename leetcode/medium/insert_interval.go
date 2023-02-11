package medium

import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)

	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	resInterval := func(a, b []int) []int {
		return []int{min2(a[0], b[0]), max2(a[1], b[1])}
	}

	overlap := func(a, b []int) bool {
		aa := min2(a[1], b[1])
		bb := max2(a[0], b[0])

		return (aa - bb) >= 0
	}

	curr := make([]int, 0, 2)
	for i := 0; i < len(intervals); i++ {
		curr = []int{intervals[i][0], intervals[i][1]}

		for i < len(intervals) && overlap(curr, intervals[i]) {
			//resInt := resInterval(intervals[i], intervals[i+1])
			//app := make([]int, len(re))
			curr = resInterval(curr, intervals[i])
			i++
		}

		i--
		res = append(res, curr)
		//intervals[i] = []int{intervals[i][0], intervals[i][1]}
	}

	//for i := 0; i < len(intervals)-1; i++ {
	//	if intervals[i][1] < newInterval[0] && intervals[i+1][0] > newInterval[1] {
	//		intervals = append(intervals, newInterval)
	//		sort.Slice(intervals, func(i, j int) bool {
	//			return intervals[i][0] < intervals[j][0]
	//		})
	//
	//		return intervals
	//	}
	//}

	//for i := 0; i < len(intervals); i++ {
	//	if len(intervals) >= i+1 && newInterval[0] >= intervals[i][0] && newInterval[1] >= intervals[i][1] && newInterval[1] < intervals[i+1][0] {
	//		res = append(res, []int{intervals[i][0], newInterval[1]})
	//		continue
	//		//for j := i+1; j < len(intervals)-1; j++ {
	//		//	in ne
	//		//}
	//	}
	//
	//	res = append(res, intervals[i])
	//}

	return res
}

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}
