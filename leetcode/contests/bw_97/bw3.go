package bw_97

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	if k == 0 {
		return n/2 + n%2
	}
	result := 0
	i, j := 0, 0
	for j < n {
		if j-i+1 <= k {
			j++
			result++
		} else {
			i++
			if j-i >= 2*k && prizePositions[j]-prizePositions[j-k]+1 <= k {
				result--
			}
		}
	}
	return result
}
