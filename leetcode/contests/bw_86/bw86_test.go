package bw_86

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubarraySum(t *testing.T) {
	assert.True(t, findSubarrays([]int{4, 2, 4}))
	assert.False(t, findSubarrays([]int{1, 2, 3, 4, 5}))
	assert.True(t, findSubarrays([]int{0, 0, 0}))
}

func TestIsStrictlyPalindromic(t *testing.T) {
	assert.False(t, isStrictlyPalindromic(9))
	assert.False(t, isStrictlyPalindromic(4))
}