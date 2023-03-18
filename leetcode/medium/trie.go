package medium

type Trie struct {
	root *TNode
	//data map[byte]*TNode
}

type TNode struct {
	Children [26]*TNode
	Value    string
}

func Constructor() Trie {
	return Trie{
		&TNode{
			Children: [26]*TNode{},
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &TNode{}
			node = node.Children[ch]
		} else {
			node = node.Children[ch]
		}
	}

	node.Value = word
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for i := 0; i < len(word); i++ {
		ch := word[i]
		if node.Children[ch-'a'] == nil {
			return false
		} else {
			node = node.Children[ch-'a']
		}
	}

	return node.Value == word
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root

	for i := 0; i < len(prefix); i++ {
		ch := prefix[i]
		if node.Children[ch-'a'] == nil {
			return false
		} else {
			node = node.Children[ch-'a']
		}
	}

	return true
}
