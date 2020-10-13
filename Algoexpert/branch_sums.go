package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	sum := 0
	res := make([]int, 0)
	branch_sums_helper(root, sum, &res)
	return res
}

func branch_sums_helper(node *BinaryTree, sum int, res *[]int) {
	if node == nil {
		return
	}
	newSum := sum + node.Value
	if node.Left == nil && node.Right == nil {
		*res = append(*res, newSum)
		return
	}
	branch_sums_helper(node.Left, newSum, res)
	branch_sums_helper(node.Right, newSum, res)
}

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	tree.Insert(values, 0)
	return tree
}

func (tree *BinaryTree) Insert(values []int, i int) *BinaryTree {
	if i >= len(values) {
		return tree
	}
	val := values[i]

	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, i+1)
	return tree
}
