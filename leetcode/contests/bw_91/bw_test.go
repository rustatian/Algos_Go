package bw_87

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBW1(t *testing.T) {
	assert.Equal(t, 5, distinctAverages([]int{9, 5, 7, 8, 7, 9, 8, 2, 0, 7}))
	assert.Equal(t, 2, distinctAverages([]int{4, 1, 4, 0, 3, 5}))
	assert.Equal(t, 1, distinctAverages([]int{1, 100}))
}

func TestBW2(t *testing.T) {
	assert.Equal(t, 89, countGoodStrings(10, 10, 2, 1))
	assert.Equal(t, 1, countGoodStrings(5, 5, 5, 2))
	assert.Equal(t, 8, countGoodStrings(3, 3, 1, 1))
	assert.Equal(t, 5, countGoodStrings(2, 3, 1, 2))
}
