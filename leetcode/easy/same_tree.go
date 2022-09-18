package easy

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil || q == nil && p != nil {
		return false
	}

	sl := make([]*TreeNode, 2)
	sl[0] = p
	sl[1] = q

	for len(sl) > 0 {
		node1 := sl[0]
		node2 := sl[1]

		sl = sl[2:]

		if node1.Val != node2.Val {
			return false
		}

		if node1.Right == nil && node2.Right != nil {
			return false
		}

		if node2.Right == nil && node1.Right != nil {
			return false
		}

		if node1.Left == nil && node2.Left != nil {
			return false
		}

		if node2.Left == nil && node1.Left != nil {
			return false
		}

		if node1.Right != nil && node2.Right != nil {
			sl = append(sl, node1.Right, node2.Right)
		}

		if node1.Left != nil && node2.Left != nil {
			sl = append(sl, node1.Left, node2.Left)
		}
	}

	return true
}
