package easy

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// search R subtree
	stack := make([]*TreeNode, 2)
	stack[0] = root
	stack[1] = root

	for len(stack) > 0 {
		n := len(stack) - 1
		node1 := stack[n]
		// rewrite
		stack = stack[:n]

		n = len(stack) - 1
		node2 := stack[n]
		stack = stack[:n]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		stack = append(stack, node1.Left)
		stack = append(stack, node2.Right)

		stack = append(stack, node1.Right)
		stack = append(stack, node2.Left)
	}

	return true
}
