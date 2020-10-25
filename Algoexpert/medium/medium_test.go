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
