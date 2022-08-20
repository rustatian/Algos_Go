package easy

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	stack := make([][]*TreeNode, 1)
	stack[0] = []*TreeNode{root1, root2}

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]

		if node[0] == nil || node[1] == nil {
			continue
		}

		node[0].Val += node[1].Val

		if node[0].Left == nil {
			node[0].Left = node[1].Left
		} else {
			stack = append(stack, []*TreeNode{node[0].Left, node[1].Left})
		}

		if node[0].Right == nil {
			node[0].Right = node[1].Right
		} else {
			stack = append(stack, []*TreeNode{node[0].Right, node[1].Right})
		}
	}

	return root1
}
