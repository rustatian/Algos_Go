package main

func TwoNumberSum(array []int, target int) []int {
	m := make(map[int]int, len(array))

	for i := 0; i < len(array); i++ {
		m[array[i]] = i
	}

	for i := 0; i < len(array); i++ {
		// hash with element and key is it's position
		num := array[i]

		if num >= 0 {
			t := target - num
			if idx, ok := m[t]; ok {
				if idx == i {
					continue
				}
				return []int{num, t}
			}
		} else {
			t := (target - num) * -1
			if idx, ok := m[t]; ok {
				if idx == i {
					continue
				}
				return []int{num, t}
			}
		}
	}
	return nil
}

func TwoNumberSum2(array []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		m := target - num // 10 - - 1 == 11
		if _, found := nums[m]; found {
			return []int{m, num}
		} else {
			nums[num] = true
		}
	}
	return nil
}
