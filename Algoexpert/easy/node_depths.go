package main

func NodeDepths(root *BinaryTree) int {
	sum := 0
	depth := 0
	node_depths_helper(root, depth, &sum)
	return sum
}

func node_depths_helper(node *BinaryTree, depth int, sum *int) {
	if node == nil {
		depth -= 1
		return
	}

	*sum += depth
	if node.Right == nil && node.Left == nil {
		return
	}

	node_depths_helper(node.Left, depth+1, sum)
	node_depths_helper(node.Right, depth+1, sum)
}
