package main

func CaesarCipherEncryptor(str string, key int) string {
	sh, offset := rune(key%26), rune(26)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if r[i] >= 'a' && r[i]+sh <= 'z' {
			r[i] += sh
		} else {
			r[i] += sh - offset
		}
	}

	return string(r)
}
