package medium

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	hm := make(map[string]*int, len(s))

	left := 0
	right := 0
	num := 0

	for right < len(s) {
		ch := string(s[right])

		if n, ok := hm[ch]; ok {
			*n++
		} else {
			hm[ch] = ptrTo(1)
		}

		for *hm[ch] > 1 {
			ch2 := string(s[left])
			if n, ok := hm[ch2]; ok {
				*n--
			}

			left++
		}

		num = int(math.Max(float64(num), float64(right-left+1)))
		right++
	}

	return num
}

func ptrTo(i int) *int {
	return &i
}
