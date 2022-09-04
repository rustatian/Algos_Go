package easy

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0, 2)
	stack = append(stack, root)

	var depth = 0
	var next = 0

	for len(stack) > 0 {
		depth++
		next = len(stack)
		for i := 0; i < next; i++ {
			//n := len(stack) - 1
			node := stack[0]
			stack = stack[:1]

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return depth
}
