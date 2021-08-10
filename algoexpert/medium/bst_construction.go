package medium

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func (tree *BST) Insert(value int) *BST {
	// no nodes
	if tree.Left == nil && tree.Right == nil && tree.Value == 0 {
		tree.Value = value
	}

	if value > tree.Value {
		if tree.Right == nil {
			tree.Right = NewBST(value)
		} else {
			return tree.Right.Insert(value)
		}

	} else {
		if tree.Left == nil {
			tree.Left = NewBST(value)
		} else {
			return tree.Left.Insert(value)
		}
	}

	return tree
}

func (tree *BST) Contains(value int) bool {
	curr := tree
	for curr != nil {
		if value == curr.Value {
			return true
		}
		if value > curr.Value {
			if curr.Right == nil {
				return false
			} else {
				curr = curr.Right
				continue
			}
		} else {
			if curr.Left == nil {
				return false
			} else {
				curr = curr.Left
				continue
			}
		}
	}
	return false
}

func (tree *BST) Remove(value int) {
	// right and leftmost
	tree.remove(value, tree)
}

func (tree *BST) remove(value int, parent *BST) {
	z := tree

	/*
	we may have 3 concepts here:
	1. We are deleting the node w/o parents -> just replace with nil
	2. Node has 1 parent
	3. Node has 2 parents
	*/
	for z != nil {
		// go to the left
		if value < z.Value {
			parent = z
			z = z.Left
		} else if value > z.Value {
			parent = z
			z = z.Right
		}


	}
}

func (tree *BST) transplant(u, parent, v *BST) {
	if parent == nil {
		panic("fdf") //tree = v
	} else if u == parent.Left {
		parent.Left = v
	} else {
		parent.Right = v
	}
	if v != nil {

	}
}

func (tree *BST) treeMinimum() *BST {
	curr := tree
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}
