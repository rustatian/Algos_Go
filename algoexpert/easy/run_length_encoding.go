package main

import "strconv"

func RunLengthEncoding(str string) string {
	res := ""
	num := 1
	for i := 1; i < len(str); i++ {
		cur := string(str[i])
		prev := string(str[i-1])

		if num == 9 {
			res += "9" + prev
			num = 1
			continue
		}

		if cur == prev {
			num++
			continue
		}

		if cur != prev {
			res += strconv.Itoa(num) + prev
			num = 1
			continue
		}
	}

	if num > 0 {
		res += strconv.Itoa(num) + string(str[len(str)-1])
	}

	return res
}
