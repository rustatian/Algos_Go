package easy

// 6.5 min

func canConstruct(ransomNote string, magazine string) bool {
	// special case with len 1
	if len(ransomNote) == 1 && len(magazine) == 1 {
		if ransomNote[0] == magazine[0] {
			return true
		}

		return false
	}

	if len(magazine) < len(ransomNote) {
		return false
	}

	// build HashMap
	magazineHm := make(map[uint8]*int, len(magazine))
	for i := 0; i < len(magazine); i++ {
		if val, ok := magazineHm[magazine[i]]; ok {
			*val += 1
		} else {
			magazineHm[magazine[i]] = ptrTo(1)
		}
	}

	for i := 0; i < len(ransomNote); i++ {
		char := ransomNote[i]

		if val, ok := magazineHm[char]; ok {
			if *val == 1 {
				delete(magazineHm, char)
				continue
			}

			*val -= 1
		} else {
			return false
		}
	}

	return true
}

func ptrTo[T any](val T) *T {
	return &val
}
