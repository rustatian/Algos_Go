package bw_95

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategorizeBox(t *testing.T) {
	//assert.Equal(t, "Heavy", categorizeBox(1000, 35, 700, 300))
	assert.Equal(t, "Both", categorizeBox(3233, 1271, 2418, 749))
}

func TestDataStream_Consec(t *testing.T) {
	ds := Constructor(4, 3)
	assert.False(t, ds.Consec(4))
	assert.False(t, ds.Consec(4))
	assert.True(t, ds.Consec(4))
	assert.False(t, ds.Consec(3))

	ds2 := Constructor(1, 2)
	assert.False(t, ds2.Consec(1))
	assert.False(t, ds2.Consec(2))
	assert.False(t, ds2.Consec(1))
	assert.True(t, ds2.Consec(1))
	assert.True(t, ds2.Consec(1))

	ds3 := Constructor(4, 1)
	assert.False(t, ds3.Consec(3))
	assert.False(t, ds3.Consec(5))
	assert.True(t, ds3.Consec(4))
	assert.True(t, ds3.Consec(4))
	assert.False(t, ds3.Consec(3))
}

func TestXORBeauty(t *testing.T) {
	assert.Equal(t, 5, xorBeauty([]int{1, 4}))
	assert.Equal(t, 34, xorBeauty([]int{15, 45, 20, 2, 34, 35, 5, 44, 32, 30}))
}
