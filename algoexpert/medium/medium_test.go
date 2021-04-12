package medium

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreeNumberSum(t *testing.T) {
	expected := [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}
	output := ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0)
	require.Equal(t, expected, output)
}

func TestSmallestDifference(t *testing.T) {
	expected := []int{28, 26}
	output := SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	require.Equal(t, expected, output)
}

func TestMoveElementToEnd(t *testing.T) {
	array := []int{5, 1, 2, 5, 5, 3, 4, 6, 7, 5, 8, 9, 10, 11, 5, 5, 12}
	toMove := 5
	expected := []int{12, 1, 2, 11, 10, 3, 4, 6, 7, 9, 8, 5, 5, 5, 5, 5, 5}
	output := MoveElementToEnd(array, toMove)
	require.Equal(t, expected, output)
}

func TestIsMonotonic(t *testing.T) {
	array := []int{-1, -5, -10, -1100, -900, -1101, -1102, -9001}
	actual := IsMonotonic(array)
	require.False(t, actual)
}

func TestSpiralTraverse(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 5},
		{9, 8, 7, 6},
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	actual := SpiralTraverse(matrix)
	require.Equal(t, expected, actual)
}

func TestLongestPeak(t *testing.T) {
	array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	actual := LongestPeak(array)
	require.Equal(t, 6, actual)
}

func TestArrayOfProducts(t *testing.T) {
	input := []int{5, 1, 4, 2}
	expected := []int{8, 40, 10, 20}
	actual := ArrayOfProducts(input)
	require.Equal(t, expected, actual)
}

func TestFirstDuplicateValue(t *testing.T) {
	input := []int{2, 1, 5, 2, 3, 3, 4}
	expected := 2
	actual := FirstDuplicateValue(input)
	require.Equal(t, expected, actual)
}

func TestBST(t *testing.T) {
	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Left = NewBST(13)
	root.Right.Left.Right = NewBST(14)
	root.Right.Right = NewBST(22)

	require.True(t, root.Contains(14))
	require.False(t, root.Contains(33))

	root.Insert(12)
	require.True(t, root.Right.Left.Left.Value == 12)

	root.Remove(10)
	require.True(t, root.Contains(10) == false)
	require.True(t, root.Value == 12)

	require.True(t, root.Contains(15))
}

func TestMergeOverlappingIntervals(t *testing.T) {
	input := [][]int{
		{4, 7},
		{6, 8},
		{1, 2},
		{3, 5},
		{9, 10},
	}
	expected := [][]int{
		{1, 2},
		{3, 8},
		{9, 10},
	}
	actual := MergeOverlappingIntervals(input)

	require.Equal(t, expected, actual)
	input2 := [][]int{
		{1, 22},
		{-20, 30},
	}
	expected2 := [][]int{
		{-20, 30},
	}
	actual2 := MergeOverlappingIntervals(input2)
	require.Equal(t, expected2, actual2)
}
