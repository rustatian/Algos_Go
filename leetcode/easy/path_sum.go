package easy

func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := make([]*TreeNode, 1)
	stack[0] = root

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[n:]

		targetSum -= node.Val

		// end
		if node.Left == nil && node.Right == nil {
			return targetSum == 0
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}

	return false
}
