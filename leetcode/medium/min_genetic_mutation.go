package medium

func minMutation(start string, end string, bank []string) int {
	q := make([]string, 0, 5)
	q = append(q, start)

	hm := make(map[string]struct{})
	hm[start] = struct{}{}

	allowedCh := []string{"A", "C", "G", "T"}

	bm := make(map[string]struct{})
	for i := 0; i < len(bank); i++ {
		bm[bank[i]] = struct{}{}
	}

	count := 0

	for len(q) > 0 {
		l := len(q)

		for j := 0; j < l; j++ {
			vertex := q[0]
			q = q[1:]

			if vertex == end {
				return count
			}

			for k := 0; k < 4; k++ {
				for kk := 0; kk < len(vertex); kk++ {
					possibleVar := vertex[:kk] + allowedCh[k] + vertex[kk+1:]

					if _, ok := hm[possibleVar]; !ok {
						if _, ok2 := bm[possibleVar]; ok2 {
							q = append(q, possibleVar)
							hm[possibleVar] = struct{}{}
						}
					}
				}
			}
		}

		count++
	}

	return -1
}
