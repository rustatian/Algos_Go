package easy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoSum(t *testing.T) {
	nums := make([]int, 0, 10)
	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 4)

	out := twoSum(nums, 6)
	require.Len(t, out, 2)
	require.Equal(t, 1, out[0])
	require.Equal(t, 2, out[1])
}

func TestTwoSum2(t *testing.T) {
	nums := make([]int, 0, 10)
	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 3)

	out := twoSum(nums, 6)
	require.Len(t, out, 2)
	require.Equal(t, 2, out[0])
	require.Equal(t, 0, out[1])
}

func TestIntersect(t *testing.T) {
	require.Equal(t, []int{2, 2}, intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	require.Equal(t, []int{4, 9}, intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
