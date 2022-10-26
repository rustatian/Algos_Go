package medium

type pair struct {
	key   string
	value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	weights := make(map[string]*pair)

	for i := 0; i < len(equations); i++ {
		eq := equations[i]
		ai := eq[0]
		bi := eq[1]
		q := values[i]

		union(&weights, ai, bi, q)
	}

	res := make([]float64, len(queries))

	for i := 0; i < len(queries); i++ {
		query := queries[i]

		ai := query[0]
		bi := query[1]

		_, ok := weights[ai]
		_, ok2 := weights[bi]

		if !ok || !ok2 {
			res[i] = -1.0
		} else {
			p1 := find(&weights, ai)
			p2 := find(&weights, bi)

			dai := p1.key
			dbi := p2.key

			daiWeight := p1.value
			dbiWeight := p2.value

			if dai != dbi {
				res[i] = -1.0
			} else {
				res[i] = daiWeight / dbiWeight
			}
		}
	}

	return res
}

func find(w *map[string]*pair, node string) *pair {
	if _, ok := (*w)[node]; !ok {
		(*w)[node] = &pair{
			key:   node,
			value: 1.0,
		}
	}

	prevPair := (*w)[node]

	if prevPair.key != node {
		newPair := find(w, prevPair.key)

		(*w)[node] = &pair{
			key:   newPair.key,
			value: newPair.value * prevPair.value,
		}
	}

	return (*w)[node]
}

func union(w *map[string]*pair, ai, bi string, q float64) {
	rootX := find(w, ai)
	rootY := find(w, bi)

	dai := rootX.key
	dbi := rootY.key

	daiWeight := rootX.value
	dbiWeight := rootY.value

	if dai != dbi {
		(*w)[dai] = &pair{
			key:   dbi,
			value: dbiWeight * q / daiWeight,
		}
	}
}
