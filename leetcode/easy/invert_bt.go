package easy

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := make([]*TreeNode, 1)
	q[0] = root

	for len(q) > 0 {
		levLen := len(q)

		for i := 0; i < levLen; i++ {
			node := q[0]

			tmp := node.Right
			node.Right = node.Left
			node.Left = tmp

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			q = q[1:]

		}
	}

	return root
}
