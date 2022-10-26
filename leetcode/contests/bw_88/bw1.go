package bw_88

import (
	"sort"
)

func equalFrequency(word string) bool {
	letters := make([]int, 26)

	for i := 0; i < len(word); i++ {
		letters[word[i]-'a']++
	}

	uniqLetters := 0
	for i := 0; i < len(letters); i++ {
		if letters[i] != 0 {
			uniqLetters++
		}
	}

	if letters[word[0]-'a'] == len(word) {
		return true
	}

	sort.Ints(letters)

	min := 0
	for i := 0; i < len(letters); i++ {
		if letters[i] > min {
			min = letters[i]
			break
		}
	}

	for i := 0; i < len(letters); i++ {
		if letters[i] == 0 {
			continue
		}

		letters[i] = letters[i] - min
	}

	nonZero := 0
	for i := 0; i < len(letters); i++ {
		if letters[i] == 0 {
			continue
		}

		if letters[i] > 0 {
			nonZero++
			continue
		}

		if letters[i] < 0 {
			return false
		}
	}

	// eq number of uniq letters
	if nonZero == 0 && uniqLetters > 1 {
		return false
	}

	if nonZero > 1 {
		return false
	}

	return true
}
