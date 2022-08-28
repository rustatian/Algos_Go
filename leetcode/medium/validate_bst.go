package medium

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 1)
	stack[0] = root

	lowStack := make([]*int, 0, 1)
	lowStack = append(lowStack, nil)
	highStack := make([]*int, 0, 1)
	highStack = append(highStack, nil)

	var low, high *int

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]

		nlow := len(lowStack) - 1
		low = lowStack[nlow]
		lowStack = lowStack[:nlow]

		nhigh := len(highStack) - 1
		high = highStack[nhigh]
		highStack = highStack[:nhigh]

		if node == nil {
			continue
		}

		val := node.Val

		if low != nil && val <= *low {
			return false
		}

		if high != nil && val >= *high {
			return false
		}

		stack = append(stack, node.Right)
		lowStack = append(lowStack, &val)
		lowStack = append(lowStack, low)

		stack = append(stack, node.Left)
		highStack = append(highStack, high)
		highStack = append(highStack, &val)

	}

	return true
}
