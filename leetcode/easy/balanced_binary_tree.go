package easy

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	res := helper(root)
	if res != -1 {
		return true
	}

	return false
}

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := helper(node.Left)
	r := helper(node.Right)

	if (l < 0 || r < 0) || abs(l-r) > 1 {
		return -1
	}

	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
