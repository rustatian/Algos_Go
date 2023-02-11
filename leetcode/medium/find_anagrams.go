package medium

import (
	"sort"
)

func findAnagrams(s string, p string) []int {
	hm := make(map[string]int)
	res := make([]int, 0, 1)
	pl := len(p)
	pp := []byte(p)

	sort.Slice(pp, func(i, j int) bool {
		return pp[i] < pp[j]
	})
	ppp := string(pp)

	for i := 0; i < len(p); i++ {
		hm[string(p[i])]++
	}

	for i := 0; i < len(s); i++ {
		if i+pl > len(s) {
			break
		}
		if matchStr([]byte(s[i:i+pl]), ppp) {
			res = append(res, i)
		}
	}

	return res
}

func matchStr(a []byte, b string) bool {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	return string(a) == b
}
