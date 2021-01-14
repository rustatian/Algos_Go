package main

func IsPalindrome(str string) bool {
	i := 0
	j := len(str) - 1

	for i <= j {
		if str[i] == str[j] {
			i++
			j--
			continue
		}
		return false
	}

	return true
}
