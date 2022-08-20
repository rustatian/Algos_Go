package medium

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	letters := [26]int{}

	for i := 0; i < len(s1); i++ {
		letters[s1[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		l2 := [26]int{}

		for j := 0; j < len(s1); j++ {
			l2[s2[i+j]-'a']++
		}

		if match(letters, l2) {
			return true
		}
	}

	return false
}

func match(ar1, ar2 [26]int) bool {
	for k := 0; k < 26; k++ {
		if ar1[k] != ar2[k] {
			return false
		}
	}
	return true
}
