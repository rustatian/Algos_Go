package easy

func minDeletionSize(strs []string) int {
	d := 0
	size := len(strs[0])

	for i := 0; i < size; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				d++
				break
			}
		}
	}

	return d
}
