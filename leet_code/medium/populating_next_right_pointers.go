package medium

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	s := make([]*Node, 1)
	s[0] = root
	level := 0

	for len(s) > 0 {
		levLen := len(s)

		nodes := s[:levLen]

		for i := 0; i < len(nodes)-1; i++ {
			nodes[i].Next = nodes[i+1]
		}

		for i := 0; i < levLen; i++ {
			node := s[0]

			if node.Left != nil {
				s = append(s, node.Left)
			}

			if node.Right != nil {
				s = append(s, node.Right)
			}

			s = s[1:]
			level++
		}
	}

	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	left := root

	for left.Left != nil {
		head := left

		for head != nil {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}

			head = head.Next
		}

		left = left.Left
	}

	return root
}
