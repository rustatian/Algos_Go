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
