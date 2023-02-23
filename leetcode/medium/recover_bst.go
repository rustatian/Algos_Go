package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	res := make([]*TreeNode, 0, 10)
	inorder(root, &res)
	hit := false

	first, second := &TreeNode{}, &TreeNode{}
	for i := 0; i < len(res)-1; i++ {
		if res[i+1].Val < res[i].Val {
			second = res[i+1]
			if !hit {
				first = res[i]
				hit = true
				continue
			}

			first.Val, second.Val = second.Val, first.Val
			break
		}
	}

	first.Val, second.Val = second.Val, first.Val
}

func inorder(node *TreeNode, res *[]*TreeNode) {
	if node == nil {
		return
	}

	inorder(node.Left, res)
	*res = append(*res, node)
	inorder(node.Right, res)
}
