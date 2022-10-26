package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostToSupplyWater(t *testing.T) {
	assert.Equal(t, 3, minCostToSupplyWater(3, []int{1, 2, 2}, [][]int{{1, 2, 1}, {2, 3, 1}}))
	assert.Equal(t, 2, minCostToSupplyWater(2, []int{1, 1}, [][]int{{1, 2, 1}, {1, 2, 2}}))
	assert.Equal(t, 2, minCostToSupplyWater(2, []int{1, 1}, [][]int{{1, 2, 1}}))
}

func TestEvaluateDivision(t *testing.T) {
	assert.Equal(t, 4, largestComponentSize([]int{4, 6, 15, 35}))
	assert.Equal(t, 2, largestComponentSize([]int{20, 50, 9, 63}))
	assert.Equal(t, 8, largestComponentSize([]int{2, 3, 6, 7, 4, 12, 21, 39}))
}
