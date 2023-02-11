package bw_97

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc1(t *testing.T) {
	assert.Equal(t, []int{1, 3, 2, 5, 8, 3, 7, 7}, separateDigits([]int{13, 25, 83, 77}))
}

func TestFunc2(t *testing.T) {
	assert.Equal(t, 2, maxCount([]int{1, 6, 5}, 5, 6))
	assert.Equal(t, 0, maxCount([]int{1, 2, 3, 4, 5, 6, 7}, 8, 1))
	assert.Equal(t, 7, maxCount([]int{11}, 7, 50))
}

func TestFunc3(t *testing.T) {
	assert.Equal(t, 7, maximizeWin([]int{1, 1, 2, 2, 3, 3, 5}, 2))
	assert.Equal(t, 2, maximizeWin([]int{1, 2, 3, 4}, 0))
	assert.Equal(t, 4, maximizeWin([]int{1, 2, 3, 4, 5}, 1))
}
