package medium

type list struct {
	col  int
	node *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0, 1)
	l := make([]*list, 1)
	l[0] = &list{
		col:  0,
		node: root,
	}

	tmp := make(map[int][]int, 10)
	mn := 0
	mx := 0

	for len(l) > 0 {
		n := l[0]
		l = l[1:]

		col := n.col

		tmp[col] = append(tmp[col], n.node.Val)

		if n.node.Left != nil {
			lst := &list{
				col:  col - 1,
				node: n.node.Left,
			}
			l = append(l, lst)

			if lst.col < mn {
				mn = lst.col
			}
		}

		if n.node.Right != nil {
			lst := &list{
				col:  col + 1,
				node: n.node.Right,
			}

			l = append(l, lst)
			if lst.col > mx {
				mx = lst.col
			}
		}
	}

	for i := mn; i <= mx; i++ {
		res = append(res, tmp[i])
	}

	return res
}
