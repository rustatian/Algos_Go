package medium

func longestPalindrome(words []string) int {
	hm := make(map[string]int, len(words))

	num := 0
	for i := 0; i < len(words); i++ {
		if _, ok := hm[words[i]]; ok {
			num += 4
			// occured once
			if hm[words[i]] == 1 {
				delete(hm, words[i])
			} else {
				hm[words[i]] -= 1
			}

			continue
		}

		hm[words[i][1:]+words[i][:1]] += 1
	}

	for k := range hm {
		// check for the patterns like aa,bb,cc..
		if k[0] == k[1] {
			num += 2
			return num
		}
	}

	return num
}
