package easy

type OrderedStream struct {
	pos    int
	values []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		pos:    1,
		values: make([]string, n+1),
	}
}

func (os *OrderedStream) Insert(idKey int, value string) []string {
	os.values[idKey] = value
	if idKey != os.pos {
		return []string{}
	}

	res := make([]string, 0, 1)

	for i := os.pos; i < len(os.values); i++ {
		if os.values[i] == "" {
			os.pos = i
			break
		}

		res = append(res, os.values[i])
	}

	return res
}
