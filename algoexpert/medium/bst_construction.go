package medium

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
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
	if value > tree.Value {
		if tree.Right == nil {
			return false
		} else if tree.Right.Value == value {
			return true
		} else {
			return tree.Right.Contains(value)
		}

	} else {
		if tree.Left == nil {
			return false
		} else if tree.Left.Value == value {
			return true
		} else {
			return tree.Left.Contains(value)
		}
	}
}

func (tree *BST) Remove(value int) *BST {
	
	return tree
}
