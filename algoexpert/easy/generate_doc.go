package main

func GenerateDocument(characters string, document string) bool {
	tmp := make(map[uint8]int, len(characters))
	for i := 0; i < len(characters); i++ {
		tmp[characters[i]] += 1
	}

	for i := 0; i < len(document); i++ {
		if _, ok := tmp[document[i]]; ok {
			if tmp[document[i]] >= 1 {
				tmp[document[i]] -= 1
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
