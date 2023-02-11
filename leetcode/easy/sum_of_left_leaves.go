package easy

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0

	q := make([]*TreeNode, 1)
	q[0] = root

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.Left == nil && node.Right == nil {
			continue
		}

		if node.Left != nil && node.Left.Right == nil && node.Left.Left == nil {
			sum += node.Left.Val
		}

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return sum
}
