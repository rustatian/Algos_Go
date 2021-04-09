package main

func FirstNonRepeatingCharacter(str string) int {
	rep := make(map[uint8]int, 26)

	for i := 0; i < len(str); i++ {
		rep[str[i]]++
	}

	for i := 0; i < len(str); i++ {
		if r, ok := rep[str[i]]; ok {
			if r == 1 {
				return i
			}
		}
	}
	return -1
}
