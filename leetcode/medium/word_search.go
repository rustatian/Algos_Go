package medium

type Vertex struct {
	next [300]*Vertex
	word string
}

func NewDataTree(word string) *Vertex {
	root := new(Vertex)

	n := root

	for j := 0; j < len(word); j++ {
		idx := word[j]

		if n.next[idx] == nil {
			n.next[idx] = &Vertex{}
		}

		n = n.next[idx]
	}

	n.word = word

	return root
}

func exist(board [][]byte, word string) bool {
	v := NewDataTree(word)
	res := make([]string, 0, 1)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, v, &res)
		}
	}

	return len(res) > 0
}

func dfs(board [][]byte, i, j int, vertex *Vertex, res *[]string) {
	ch := board[i][j]

	if ch == '#' || vertex.next[ch] == nil {
		return
	}

	vertex = vertex.next[ch]

	if vertex.word != "" {
		*res = append(*res, vertex.word)
		vertex.word = ""
	}

	board[i][j] = '#'

	// directions
	if i > 0 {
		dfs(board, i-1, j, vertex, res)
	}
	if j > 0 {
		dfs(board, i, j-1, vertex, res)
	}
	if i < len(board)-1 {
		dfs(board, i+1, j, vertex, res)
	}
	if j < len(board[0])-1 {
		dfs(board, i, j+1, vertex, res)
	}

	board[i][j] = ch
}
