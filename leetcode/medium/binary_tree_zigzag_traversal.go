package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0, 10)
	q := make([]*TreeNode, 1)
	q[0] = root

	level := 0
	mode := true

	for len(q) > 0 {
		levLen := len(q)

		res = append(res, []int{})
		res[level] = make([]int, levLen)

		idx := 0
		for i := 0; i < levLen; i++ {
			var node *TreeNode

			node = q[0]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			switch mode {
			case true:
				res[level][idx] = node.Val
			case false:
				res[level][len(res[level])-1-idx] = node.Val
			}

			q = q[1:]
			idx++
		}

		if mode {
			mode = false
		} else {
			mode = true
		}

		level++
	}

	return res
}
