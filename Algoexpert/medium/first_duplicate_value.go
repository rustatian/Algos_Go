package medium

// O(n) time, O(n) space (map)
func FirstDuplicateValue2(array []int) int {
	m := make(map[int]int)
	for i := 0; i < len(array); i++ {
		if _, ok := m[array[i]]; ok {
			return array[i]
		}
		m[array[i]] = 1
	}
	return -1
}

// O(n) time, O(1) space
func FirstDuplicateValue(array []int) int {
	for i := 0; i < len(array); i++ {
		index := abs(array[i])
		if (array[index-1]) < 0 {
			return index
		}
		array[index-1] *= -1
	}
	return -1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
