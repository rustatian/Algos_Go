package bw_133

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 3, minimumOperations([]int{1, 2, 3, 4}))
	assert.Equal(t, 0, minimumOperations([]int{3, 6, 9}))
}
