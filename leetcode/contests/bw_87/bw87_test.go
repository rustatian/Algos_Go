package bw_87

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBW1(t *testing.T) {
	//assert.Equal(t, 3, countDaysTogether("08-15", "08-18", "08-16", "08-19"))
	//assert.Equal(t, 0, countDaysTogether("10-01", "10-31", "11-01", "12-31"))
	assert.Equal(t, 0, countDaysTogether("06-01", "07-31", "04-01", "12-31"))
}

func TestBW2(t *testing.T) {
	assert.Equal(t, 2, matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8}))
	assert.Equal(t, 1, matchPlayersAndTrainers([]int{1, 1, 1}, []int{10}))
}

func TestBW3(t *testing.T) {
	assert.Equal(t, []int{4, 3, 3, 2, 3, 4, 3, 2, 1}, smallestSubarrays([]int{9, 5, 0, 10, 7, 2, 9, 2, 4}))
	assert.Equal(t, []int{4, 3, 2, 2, 1, 1}, smallestSubarrays([]int{4, 0, 5, 6, 3, 2}))
	assert.Equal(t, []int{2, 1, 1}, smallestSubarrays([]int{8, 10, 8}))
	assert.Equal(t, []int{3, 3, 2, 2, 1}, smallestSubarrays([]int{1, 0, 2, 1, 3}))
	assert.Equal(t, []int{2, 1}, smallestSubarrays([]int{1, 2}))
}

func TestBW4(t *testing.T) {

}
