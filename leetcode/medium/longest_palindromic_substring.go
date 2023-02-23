package medium

// https://en.wikipedia.org/wiki/Longest_palindromic_substring
func longestPalindrome2(s string) string {
	bogusS := "|"
	for i := 0; i < len(s); i++ {
		bogusS += string(s[i])
		bogusS += "|"
	}

	PalindromeRadii := make([]int, len(s)*2+1)
	center := 0
	longestBogus := 0

	for center < len(bogusS) {
		radius := 0

		for center-(radius+1) >= 0 &&
			center+radius+1 < len(bogusS) &&
			bogusS[center-radius+1] == bogusS[center+radius+1] {

			radius += 1
		}

		val := 2*radius

		if val > longestBogus {
			longestBogus = val
		}
		PalindromeRadii[center] = val
		center += 1
	}

	longestInS := (longestBogus-1)/2
	print(longestInS)

	return ""
}
