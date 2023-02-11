package easy

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diff := 10000000

	inorder(root, root.Val, &diff)

	return diff
}

func inorder(node *TreeNode, val int, diff *int) {
	if node == nil {
		return
	}

	if val != node.Val {
		a := abs1(val, node.Val)
		if *diff > a {
			*diff = a
		}
	}

	inorder(node.Left, node.Val, diff)
	inorder(node.Right, node.Val, diff)
}

func abs1(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}
