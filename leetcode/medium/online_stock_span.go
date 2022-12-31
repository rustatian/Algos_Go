package medium

type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{{1<<31 - 1, 0}}}
}

func (st *StockSpanner) Next(price int) int {
	res := 1

	for len(st.stack) > 0 && st.stack[len(st.stack)-1][0] <= price {
		res += st.stack[len(st.stack)-1][1]
		st.stack = st.stack[:len(st.stack)-1]
	}

	st.stack = append(st.stack, [2]int{price, res})

	return res
}
