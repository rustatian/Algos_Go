package bw_96

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommon(t *testing.T) {
	assert.Equal(t, 15, getCommon([]int{11, 15, 28, 31, 36, 42, 46, 54, 58, 63, 64, 67, 75, 76, 76, 79, 83, 85, 87, 90},
		[]int{3, 6, 8, 13, 15, 19, 22, 25, 29, 29, 32, 35, 43, 43, 48, 55, 81, 90, 91, 94}))
	assert.Equal(t, 2, getCommon([]int{1, 2, 3}, []int{2, 4}))
	assert.Equal(t, 2, getCommon([]int{1, 2, 3, 6}, []int{2, 3, 4, 5}))
}

func TestMinOperations(t *testing.T) {
	assert.Equal(t, int64(0), minOperations([]int{0, 0, 0}, []int{0, 0, 0}, 0))
	assert.Equal(t, int64(-1), minOperations([]int{5, 1, 0}, []int{9, 7, 6}, 2))
	assert.Equal(t, int64(-1), minOperations([]int{4, 3, 1, 4}, []int{2, 3, 6, 1}, 10))
	assert.Equal(t, int64(6), minOperations([]int{13, 6, 10, 16}, []int{1, 16, 12, 16}, 2))
	assert.Equal(t, int64(2), minOperations([]int{4, 3, 1, 4}, []int{1, 3, 7, 1}, 3))
	assert.Equal(t, int64(-1), minOperations([]int{3, 8, 5, 2}, []int{2, 4, 1, 6}, 1))
}

func TestMaxSubsq(t *testing.T) {
	assert.Equal(t, int64(168), maxScore([]int{2,1,14,12}, []int{11,7,13,6}, 3))
	//assert.Equal(t, int64(12), maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
	//assert.Equal(t, int64(30), maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1))
}
