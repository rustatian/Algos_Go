package medium

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}

		return root
	}

	tn := root

	for {
		if tn.Val > val {
			if tn.Left == nil {
				tn.Left = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}

				return root
			}

			tn = tn.Left
		}

		if tn.Val < val {
			if tn.Right == nil {
				tn.Right = &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				return root
			}

			tn = tn.Right
		}
	}
}
