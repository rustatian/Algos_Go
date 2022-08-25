package medium

func combine(n int, k int) [][]int {
	res := make([][]int, 0, 2)
	curr := make([]int, 0, 2)
	backtrack2(1, n, k, &curr, &res)
	return res
}

func backtrack2(first int, n, k int, curr *[]int, res *[][]int) {
	if len(*curr) == k {
		tmp := make([]int, len(*curr))
		copy(tmp, *curr)
		*res = append(*res, tmp[:])
	}

	for i := first; i < n+1; i++ {
		*curr = append(*curr, i)

		backtrack2(i+1, n, k, curr, res)

		*curr = (*curr)[:len(*curr)-1]
	}
}
