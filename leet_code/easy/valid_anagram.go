package easy

// 5m

import (
	"sort"
)

type S []byte

func (s S) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s S) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s S) Len() int {
	return len(s)
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss := S(s)
	sort.Sort(ss)

	tt := S(t)
	sort.Sort(tt)

	for i := 0; i < len(ss); i++ {
		if ss[i] != tt[i] {
			return false
		}
	}

	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make([]int, 26)

	for i := 0; i < len(s); i++ {
		letters[s[i]-'a']++
		letters[t[i]-'a']--
	}

	for i := 0; i < len(letters); i++ {
		if letters[i] != 0 {
			return false
		}
	}

	return true
}
