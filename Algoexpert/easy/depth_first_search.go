package main

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, ch := range n.Children {
		array = ch.DepthFirstSearch(array)
	}

	return array
}


