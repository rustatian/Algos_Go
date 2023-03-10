package easy

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	arr := make([]int, n+1)
	arr[1] = 1
	arr[2] = 2

	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}
