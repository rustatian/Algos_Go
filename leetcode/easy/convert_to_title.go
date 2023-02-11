package easy

func convertToTitle(columnNumber int) string {
	table := map[int]string{
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
		14: "N",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
		26: "Z",
	}

	if columnNumber <= 26 {
		return table[columnNumber]
	}

	res := ""
	tmpSl := make([]string, 0, 10)

	for columnNumber > 0 {
		columnNumber -= 1
		tmp := columnNumber % 26
		tmpSl = append(tmpSl, string(rune(tmp+'A')))
		columnNumber /= 26
	}

	for i := len(tmpSl) - 1; i >= 0; i-- {
		res += tmpSl[i]
	}

	return res
}
