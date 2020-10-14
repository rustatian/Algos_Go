package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}

func TestNode_DepthFirstSearch(t *testing.T) {
	var graph = NewNode("A").AddChildren("B", "C", "D")
	graph.Children[0].AddChildren("E").AddChildren("F")
	graph.Children[2].AddChildren("G").AddChildren("H")
	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
	graph.Children[2].Children[0].AddChildren("K")
	output := graph.DepthFirstSearch([]string{})
	expected := []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"}
	require.Equal(t, expected, output)
}

func TestBranchSums(t *testing.T) {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	expected := []int{15, 16, 18, 10, 11}
	output := BranchSums(tree)
	require.Equal(t, expected, output)
}
func TestIsValidSubsequence(t *testing.T) {
	assert.True(t, IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{22, 25, 6}))
}

func TestTwoNumberSum(t *testing.T) {
	sl := []int{-1, 11}
	assert.Equal(t, TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10), sl)
}

func TestNode_Depth(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Left.Left = &BinaryTree{Value: 8}
	root.Left.Left.Right = &BinaryTree{Value: 9}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right = &BinaryTree{Value: 3}
	root.Right.Left = &BinaryTree{Value: 6}
	root.Right.Right = &BinaryTree{Value: 7}
	actual := NodeDepths(root)
	require.Equal(t, 16, actual)
}