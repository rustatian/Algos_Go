package medium

type q struct {
	node *TreeNode
	sum  int
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	stack := make([]*q, 1)
	stack[0] = &q{
		node: root,
		sum:  0,
	}

	for len(stack) > 0 {
		qq := stack[0]
		stack = stack[1:]

		node := qq.node
		sum := qq.sum

		if node != nil {
			num := sum*10 + node.Val // tens, hundreds, thousands.. (*10)

			if node.Left == nil && node.Right == nil {
				res += num
				continue
			}

			if node.Right != nil {
				stack = append(stack, &q{
					node: node.Right,
					sum:  num,
				})
			}

			if node.Left != nil {
				stack = append(stack, &q{
					node: node.Left,
					sum:  num,
				})
			}
		}

	}

	return res
}
