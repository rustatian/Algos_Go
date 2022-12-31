package easy

func makeGood(s string) string {
	sl := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		// uppercase
		if len(sl) > 0 && sl[len(sl)-1]^s[i] == 32 {
			sl = sl[:len(sl)-1]
			continue
		}

		sl = append(sl, s[i])
		ff := sl[len(sl)-1]^s[i]
		_ = ff
	}

	return string(sl)
}
