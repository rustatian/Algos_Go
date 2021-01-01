package main

func GetNthFib(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	// just init
	fibo := 1
	prevFibo := 1

	for i := 2; i < n; i++ {
		tmp := fibo
		fibo += prevFibo
		prevFibo = tmp
	}

	return prevFibo
}
