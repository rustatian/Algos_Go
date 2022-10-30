package easy

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	ss := []byte(s)

	st := 0
	end := len(ss) - 1

	for st < end {
		if ss[st] == ' ' {
			st++
			continue
		}

		if ss[end] == ' ' {
			end--
			continue
		}

		if !(ss[st] >= 'a' && ss[st] <= 'z') && !(ss[st] >= 'A' && ss[st] <= 'Z') {
			st++
			continue
		}

		if !(ss[end] >= 'a' && ss[end] <= 'z') && !(ss[end] >= 'A' && ss[end] <= 'Z') {
			end--
			continue
		}

		if ss[st] >= 'A' && ss[st] <= 'Z' {
			ss[st] = ss[st] + 32
		}

		if ss[end] >= 'A' && ss[end] <= 'Z' {
			ss[end] = ss[end] + 32
		}

		if ss[st] == ss[end] {
			st++
			end--
			continue
		}

		return false
	}

	return true
}
