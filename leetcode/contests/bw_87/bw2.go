package bw_87

import (
	"sort"
)

func matchPlayersAndTrainers(players []int, trainers []int) int {
	if len(players) == 1 && len(trainers) == 1 {
		if players[0] <= trainers[0] {
			return 1
		}

		return 0
	}

	count := 0

	sort.Ints(players)
	sort.Ints(trainers)

	st1 := 0
	st2 := 0

	end1 := len(players) - 1
	end2 := len(trainers) - 1

	for {
		if st1 > end1 || st2 > end2 {
			break
		}

		if players[st1] <= trainers[st2] {
			count++
			st1++
			st2++
			continue
		}

		if players[st1] > trainers[st2] {
			st2++
			continue
		}
	}

	return count
}
