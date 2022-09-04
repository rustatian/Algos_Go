package main

import (
	"math"
)

// A ^ 2xM = (A^M) ^ 2
// A^(MxN) = A^M x A^N
// That code compared just to 7x7x7x7x7x7x7 <- computations like this

func power(a, p float64) float64 {
	if p == 1 {
		return a
	}

	if p == 0 {
		return 1
	}

	var i float64 = 2
	var res float64 = 1

	for i <= p {
		res = math.Pow(a, i)
		i = i * 2
	}
	//If we have power for example 35 -> in prev step we calculate only power of 32 and must add power of 3
	if (p - i) != 0 {
		res = math.Pow(a, p-i/2) * res
	}

	return res
}
