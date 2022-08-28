package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	q := make([]*TreeNode, 1)
	q[0] = root
	level := 0

	for len(q) > 0 {
		levLen := len(q)
		res = append(res, []int{})

		for i := 0; i < levLen; i++ {
			node := q[0]
			res[level] = append(res[level], node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			q = q[1:]

		}
		level++
	}

	return res
}
