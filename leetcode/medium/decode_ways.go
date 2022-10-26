package medium

var globalMap = map[string]int{
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"11": 11,
	"12": 12,
	"13": 13,
	"14": 14,
	"15": 15,
	"16": 16,
	"17": 17,
	"18": 18,
	"19": 19,
	"20": 20,
	"21": 21,
	"22": 22,
	"23": 23,
	"24": 24,
	"25": 25,
	"26": 26,
}

func numDecodings(s string) int {
	q := make([]int, len(s)+1)
	q[0] = 1

	if globalMap[s[:1]] == 0 {
		q[1] = 0
	} else {
		q[1] = 1
	}

	for i := 2; i < len(q); i++ {
		if globalMap[s[i-1:i]] != 0 {
			q[i] = q[i-1]
		}

		if globalMap[s[i-2:i]] >= 10 && globalMap[s[i-2:i]] <= 26 {
			q[i] += q[i-2]
		}
	}

	return q[len(s)]
}
