package design_add_search_words

type TNode struct {
	Children   [26]*TNode
	IsTerminal bool
	Value      string
}

type WordDictionary struct {
	root *TNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		&TNode{
			Children: [26]*TNode{},
		},
	}
}

func (wd *WordDictionary) AddWord(word string) {
	node := wd.root
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &TNode{}
			node = node.Children[ch]
		} else {
			node = node.Children[ch]
		}
	}

	node.IsTerminal = true
	node.Value = word
}

func (wd *WordDictionary) Search(word string) bool {
	node := wd.root
	return wd.search(word, node)
}

func (wd *WordDictionary) search(word string, node *TNode) bool {
	for i := 0; i < len(word); i++ {
		ch := word[i]

		switch ch {
		case '.':
			// search in every node
			allNil := true
			for j := 0; j < 26; j++ {
				if node.Children[j] == nil {
					continue
				}

				allNil = false
				child := node.Children[j]
				if wd.search(word[i+1:], child) {
					return true
				}
			}

			if allNil {
				return false
			}

			return false

		default:
			if node.Children[ch-'a'] == nil {
				return false
			} else {
				node = node.Children[ch-'a']
			}
		}
	}

	return node.IsTerminal
}
