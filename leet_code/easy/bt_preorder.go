package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 1)

	n := root

	for n != nil {
		if n.Left == nil {
			res = append(res, n.Val)
			n = n.Right
		} else {
			predecessor := n.Left

			for predecessor.Right != nil && predecessor.Right != n {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				res = append(res, n.Val)
				predecessor.Right = n
				n = n.Left
			} else {
				predecessor.Right = nil
				n = n.Right
			}
		}
	}

	return res
}

func printPreOrd(node *TreeNode, elm *[]int) {
	if node == nil {
		return
	}

	*elm = append(*elm, node.Val)

	printPreOrd(node.Left, elm)
	printPreOrd(node.Right, elm)
}
