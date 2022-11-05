package hard

type Vertex struct {
	next [26]*Vertex
	word string
}

func NewDataTree(words []string) *Vertex {
	root := new(Vertex)

	for i := 0; i < len(words); i++ {
		n := root

		for j := 0; j < len(words[i]); j++ {
			idx := words[i][j] - 'a'

			if n.next[idx] == nil {
				n.next[idx] = &Vertex{}
			}

			n = n.next[idx]
		}
		n.word = words[i]
	}

	return root
}

func findWords(board [][]byte, words []string) []string {
	v := NewDataTree(words)
	res := make([]string, 0, 1)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, v, &res)
		}
	}

	return res
}

func dfs(board [][]byte, i, j int, vertex *Vertex, res *[]string) {
	ch := board[i][j]

	if ch == '#' || vertex.next[ch-'a'] == nil {
		return
	}

	vertex = vertex.next[ch-'a']

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
	if i < len(board) - 1 {
		dfs(board, i+1, j, vertex, res)
	}
	if j < len(board[0]) - 1 {
		dfs(board, i, j+1, vertex, res)
	}

	board[i][j] = ch
}
