package main

import "sort"

func NonConstructableChange(coins []int) int {
	if len(coins) == 0 {
		return 1
	}

	if len(coins) == 1 && coins[0] > 1 {
		return 1
	}

	if len(coins) == 1 && coins[0] == 1 {
		return 2
	}

	// Write your code here.
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})

	var sum int
	for i := 0; i < len(coins)-1; i++ {
		sum += coins[i]

		if coins[i+1] > sum+1 {
			return sum + 1
		}
		if coins[i] < sum+1 {
			continue
		}
	}

	return sum + coins[len(coins)-1] + 1
}
