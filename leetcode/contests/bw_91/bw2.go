package bw_87

func countGoodStrings(low int, high int, zero int, one int) int {
	mod := 1000000007

	res := 0
	if (zero == one) && zero == 1 {
		for i := low; i <= high; i++ {
			res += calc(i)
		}

		return res % mod
	}

	//mod := 1e9+7

	q := make([]int, high+1)
	q[0] = 1

	for i := 1; i <= high; i++ {
		if i >= zero {
			q[i] = (q[i] + q[i-zero]) % mod
		}
		if i >= one {
			q[i] = (q[i] + q[i-one]) % mod
		}
		if i >= low {
			res = (res + q[i]) % mod
		}
	}

	return res
}

func calc(n int) int {
	return 2 + 2*(n-1) + (n-2)*(n-1)
}
