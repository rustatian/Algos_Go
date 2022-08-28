package easy

func findTarget(root *TreeNode, k int) bool {
	hm := make(map[int]int)

	nums := make([]int, 0, 1)

	bst := root
	stack := make([]*TreeNode, 1)
	stack[0] = bst

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]

		nums = append(nums, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	for i := 0; i < len(nums); i++ {
		diff := k - nums[i]

		if _, ok := hm[diff]; ok {
			return true
		}

		hm[nums[i]] = i
	}

	return false
}
