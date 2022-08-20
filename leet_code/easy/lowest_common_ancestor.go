package easy

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	bst := root
	stack := make([]*TreeNode, 1)
	stack[0] = bst

	pVal := p.Val
	qVal := q.Val

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]

		nVal := node.Val

		if pVal > nVal && qVal > nVal {
			stack = append(stack, node.Right)
		} else if pVal < nVal && qVal < nVal {
			stack = append(stack, node.Left)
		} else {
			return node
		}
	}

	return nil
}
