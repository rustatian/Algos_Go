package bw_88

type LUPrefix struct {
	s      []int
	prevUp int
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		s: make([]int, n+1),
	}
}

func (l *LUPrefix) Upload(video int) {
	l.s[video] = 1
	if l.prevUp > video {
		return
	}
	l.prevUp = video
}

func (l *LUPrefix) Longest() int {
	if l.s[1] == 0 {
		return 0
	}

	longest := 0
	for i := 0; i < l.prevUp; i++ {
		if l.s[i] == 1 {
			longest++
		} else {
			break
		}
	}

	return longest
}
