package medium

type list2 struct {
	col  int
	node *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := make([]*list2, 1)
	l[0] = &list2{
		col:  0,
		node: root,
	}

	mn := 0
	mx := 0

	for len(l) > 0 {
		n := l[0]
		l = l[1:]

		col := n.col

		if n.node.Left != nil {
			lst := &list2{
				col:  col - 1,
				node: n.node.Left,
			}
			l = append(l, lst)

			if lst.col < mn {
				mn = lst.col
			}
		}

		if n.node.Right != nil {
			lst := &list2{
				col:  col + 1,
				node: n.node.Right,
			}

			l = append(l, lst)
			if lst.col > mx {
				mx = lst.col
			}
		}
	}

	if mn < 0 {
		mn = mn * -1
	}
	if mx < 0 {
		mx = mx * -1
	}

	return mn + mx
}
