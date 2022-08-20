package easy

func searchBST(root *TreeNode, val int) *TreeNode {
	st := make([]*TreeNode, 1)
	st[0] = root

	for len(st) > 0 {
		n := len(st) - 1
		node := st[n]
		st = st[1:]

		if node.Val == val {
			return node
		}

		if node.Val > val {
			if node.Left == nil {
				return nil
			}
			st = append(st, node.Left)
		}

		if node.Val < val {
			if node.Right == nil {
				return nil
			}
			st = append(st, node.Right)
		}
	}

	return nil
}
