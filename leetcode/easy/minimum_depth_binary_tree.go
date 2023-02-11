package easy

type list struct {
	node  *TreeNode
	depth int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*list{{
		node:  root,
		depth: 1,
	}}

	var res = 0

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		res = node.depth

		if node.node.Left == nil && node.node.Right == nil {
			break
		}

		if node.node.Left != nil {
			stack = append(stack, &list{
				node:  node.node.Left,
				depth: res + 1,
			})
		}
		if node.node.Right != nil {
			stack = append(stack, &list{
				node:  node.node.Right,
				depth: res + 1,
			})
		}
	}

	return res
}