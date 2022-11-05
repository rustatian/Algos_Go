package easy

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}

	count := make([]uint8, 128)
	num := 0

	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	for i := 0; i < len(count); i++ {
		num += int(count[i] / 2 * 2)

		if num%2 == 0 && count[i]%2 == 1 {
			num++
		}
	}

	return num
}
