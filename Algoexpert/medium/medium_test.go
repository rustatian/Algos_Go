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
