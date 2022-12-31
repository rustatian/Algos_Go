package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedian(t *testing.T) {
	c := Constructor()
	c.AddNum(1)
	c.AddNum(2)
	assert.Equal(t, 1.5, c.FindMedian())
	c.AddNum(3)
	assert.Equal(t, 2.0, c.FindMedian())
}

func TestWordSearch(t *testing.T) {
	assert.Equal(t, nil, findWords([][]byte{{'a'}}, nil))
	assert.Equal(t, []string{"eat", "oath"}, findWords([][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}))
}

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
