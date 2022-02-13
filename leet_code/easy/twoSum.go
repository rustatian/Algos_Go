package main

// array, hashtable

func twoSum2(nums []int, target int) []int {
	lp := 0
	rp := len(nums) - 1

	for {
		res := nums[lp] + nums[rp]
		switch {
		case res == target:
			return []int{lp + 1, rp + 1}
		case res < target:
			lp = lp + 1
		case res > target:
			rp = rp - 1
		}

		if rp == lp {
			return nil
		}
	}
}

func twoSum(nums []int, target int) []int {
	// num - orig index
	hm := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if idx, ok := hm[diff]; ok {
			return []int{idx, i}
		}

		hm[nums[i]] = i
	}

	return nil
}
