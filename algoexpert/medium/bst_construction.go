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
	for z != nil {
		if z.Value > value {
			// save the parent node
			parent = z
			z = z.Left
		} else if z.Value < value {
			// save the parent node
			parent = z
			z = z.Right
		} else {
			if z.Left != nil && z.Right != nil {
				z = tree.treeMinimum()
			}
			// found
			// invariant 1
			/*
						q
						|
						Z
					   / \
					  /   \
				    NIL    R
				          / \
				         /   \

				We have Z to delete
				Z has no left child
				Z has right child
							q
							|
							R
						   / \
						  /   \
			*/
			if z.Left == nil {
				tree.transplant(z, parent, z.Right)
			}

			// invariant 2 piashchynski_valery@hotmail.com +375292704404
			/*
							q
							|
							Z
						   / \
						  /   \
					     L    NIL
						/ \
				       /   \
					We have Z to delete
					Z has no right child
					Z has left child
								q
								|
								R
							   / \
							  /   \
			*/

			if z.Right == nil {
				tree.transplant(z, parent, z.Left)
			}

			// both nodes
			if z.Left != nil && z.Right != nil {

			}
		}
	}
}

func (tree *BST) transplant(u, parent, v *BST) {
	if parent == nil {
		panic("fdf")//tree = v
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
	for curr != nil {
		curr = curr.Left
	}
	return curr
}
