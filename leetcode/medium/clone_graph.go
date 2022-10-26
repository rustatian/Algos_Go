package medium

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node2 struct {
	Val       int
	Neighbors []*Node2
}

// initial node
var hm = make(map[*Node2]*Node2, 1)

func cloneGraph(node *Node2) *Node2 {
	if node == nil {
		return node
	}

	if n, ok := hm[node]; ok {
		return n
	}

	nn := new(Node2)
	nn.Val = node.Val
	nn.Neighbors = make([]*Node2, 0, 1)

	hm[node] = nn

	for i := 0; i < len(node.Neighbors); i++ {
		nn.Neighbors = append(nn.Neighbors, cloneGraph(node.Neighbors[i]))
	}

	return nn
}
