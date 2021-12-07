package main

// TournamentWinner O(n) - num of competitions + o(k) - num of teams
func TournamentWinner(competitions [][]string, results []int) string {
	// key - team, val - number of wins
	ret := make(map[string]int, len(results))
	bestTeam := ""
	points := 0

	// for the every results
	for i := 0; i < len(results); i++ {
		// [home, away]
		team := competitions[i]

		var won string
		switch results[i] {
		// home team wins, index 0
		case 1:
			won = team[0]
			// away team wins, index 1
		case 0:
			won = team[1]
		}

		if _, ok := ret[won]; ok {
			val := ret[won]
			val += 3
			ret[won] = val

			if points < val {
				points = val
				bestTeam = won
			}
		} else {
			val := 3

			if points < val {
				points = val
				bestTeam = won
			}

			ret[won] = 3
		}

	}

	return bestTeam
}
