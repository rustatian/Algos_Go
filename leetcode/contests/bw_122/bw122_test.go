package bw_122

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 6, minimumCost([]int{1, 2, 3, 12}))
	assert.Equal(t, 12, minimumCost([]int{5, 3, 4}))
	assert.Equal(t, 12, minimumCost([]int{10, 3, 1, 1}))
}

func Test2(t *testing.T) {
	assert.True(t, canSortArray([]int{1, 2, 3, 4, 5}))
	assert.True(t, canSortArray([]int{8, 4, 2, 30, 15}))
	assert.False(t, canSortArray([]int{3, 16, 8, 4, 2}))
}
