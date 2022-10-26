package medium

const MOD = 1000000007

func numRollsToTarget(n int, k int, target int) int {
	memo := make([][]int, n+1)

	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, target+1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}

	return solve(memo, 0, n, 0, target, k)
}

func solve(memo [][]int, diceIdx, n, currSum, target, k int) int {
	if diceIdx == n {
		if currSum == target {
			return 1
		}
		return 0
	}

	if memo[diceIdx][currSum] != -1 {
		return memo[diceIdx][currSum]
	}

	ways := 0

	for i := 1; i <= min1(k, target-currSum); i++ {
		ways = (ways + solve(memo, diceIdx+1, n, currSum+i, target, k)) % MOD
	}

	memo[diceIdx][currSum] = ways

	return memo[diceIdx][currSum]
}

func min1(a, b int) int {
	if a < b {
		return a
	}

	return b
}
