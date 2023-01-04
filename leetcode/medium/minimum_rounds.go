package medium

func minimumRounds(tasks []int) int {
	hm := make(map[int]int, len(tasks))

	for i := 0; i < len(tasks); i++ {
		// tasks[i] - num
		if _, ok := hm[tasks[i]]; !ok {
			hm[tasks[i]] = 1
			continue
		}

		hm[tasks[i]]++
	}

	steps := 0

	for _, v := range hm {
		// it's not possible
		if v == 1 {
			return -1
		}

		if v == 2 {
			steps++
			continue
		}

		// we can just use max steps
		num := v % 3
		if num == 0 {
			steps += v / 3
			continue
		}

		res := v/3 + 1
		steps += res
	}

	return steps
}
