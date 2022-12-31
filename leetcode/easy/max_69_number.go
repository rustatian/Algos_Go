package easy

func maximum69Number(num int) int {
	n := num
	res := num

	mul := 1

	for num > 0 {
		if num%10 == 6 {
			res = n + (mul * 3)
		}
		mul *= 10
		num /= 10
	}

	return res
}
