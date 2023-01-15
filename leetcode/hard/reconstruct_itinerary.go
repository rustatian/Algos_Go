package hard

import (
	"sort"
)

var flights = map[string][]string{}
var visits = map[string]struct{}{}
var num = 0
var result []string

func findItinerary(tickets [][]string) []string {
	for i := 0; i < len(tickets); i++ {
		origin := tickets[i][0]
		dest := tickets[i][1]

		if dests, ok := flights[origin]; ok {
			dests = append(dests, dest)
		} else {
			dests := make([]string, 1)
			dests[0] = dest
			flights[origin] = dests
		}

		for k, v := range flights {
			sort.Slice(v, func(i, j int) bool {

			})
		}

	}

	return result
}
