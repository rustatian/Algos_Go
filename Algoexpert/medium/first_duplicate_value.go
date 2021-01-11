package medium

func FirstDuplicateValue(array []int) int {
	m := make(map[int]int)
	for i := 0; i < len(array); i++ {
		if _, ok := m[array[i]]; ok {
			return array[i]
		}
		m[array[i]] = 1
	}
	return -1
}
