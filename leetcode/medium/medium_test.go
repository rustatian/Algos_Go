package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLeadsToDestination(t *testing.T) {
	assert.False(t, leadsToDestination(3, [][]int{{0, 1}, {0, 2}}, 0, 2))
}

func TestContinuousSum(t *testing.T) {
	assert.True(t, checkSubarraySum([]int{1, 3, 6, 0, 9, 6, 9}, 7))
	assert.True(t, checkSubarraySum([]int{1, 2, 3}, 6))
	assert.False(t, checkSubarraySum([]int{5, 2, 4}, 5))
	assert.False(t, checkSubarraySum([]int{0, 1, 0, 3, 0, 4, 0, 4, 0}, 5))
	assert.True(t, checkSubarraySum([]int{5, 0, 0, 0}, 3))
	assert.True(t, checkSubarraySum([]int{23, 2, 4, 6, 6}, 7))
	assert.True(t, checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	assert.True(t, checkSubarraySum([]int{23, 2, 6, 4, 7}, 6))
	assert.False(t, checkSubarraySum([]int{23, 2, 6, 4, 7}, 13))
}

func TestCloneGraph(t *testing.T) {
	node := new(Node2)

	node.Val = 2
	node.Neighbors = append(node.Neighbors)
}

func TestAllPathsFromSourceToTarget(t *testing.T) {
	assert.Equal(t, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}, allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
	assert.Equal(t, [][]int{{0, 1, 3}, {0, 2, 3}}, allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}

func TestNumRollsToTarget(t *testing.T) {
	assert.Equal(t, 1, numRollsToTarget(1, 6, 3))
	assert.Equal(t, 6, numRollsToTarget(2, 6, 7))
	assert.Equal(t, 1000000007, numRollsToTarget(30, 30, 500))
}

func TestDecodeWays(t *testing.T) {
	assert.Equal(t, 1, numDecodings("10"))
	assert.Equal(t, 2, numDecodings("12"))
	assert.Equal(t, 3, numDecodings("226"))
	assert.Equal(t, 0, numDecodings("06"))
}

func TestEvaluateDivision(t *testing.T) {
	assert.Equal(t, []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}, calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
}

func TestSmallestStringWithSwaps(t *testing.T) {
	assert.Equal(t, "bdfgjjnuvrx", smallestStringWithSwaps("vbjjxgdfnru", [][]int{{8, 6}, {3, 4}, {5, 2}, {5, 5}, {3, 5}, {7, 10}, {6, 0}, {10, 0}, {7, 1}, {4, 8}, {6, 2}}))
	assert.Equal(t, "bacd", smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}))
	assert.Equal(t, "abcd", smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
	assert.Equal(t, "abc", smallestStringWithSwaps("cba", [][]int{{0, 1}, {1, 2}}))
}

func TestEarliestAcq(t *testing.T) {
	assert.Equal(t, 20190301, earliestAcq([][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}, 6))
	assert.Equal(t, 3, earliestAcq([][]int{{0, 2, 0}, {1, 0, 1}, {3, 0, 3}, {4, 1, 2}, {7, 3, 1}}, 4))
}

func TestNumberOfComponents(t *testing.T) {
	assert.Equal(t, 1, countComponents(4, [][]int{{0, 1}, {2, 3}, {1, 2}}))
	assert.Equal(t, 2, countComponents(5, [][]int{{0, 1}, {0, 4}, {1, 4}, {2, 3}}))
	assert.Equal(t, 2, countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}))
	assert.Equal(t, 1, countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}))
}

func TestValidTree(t *testing.T) {
	assert.False(t, validTree(5, [][]int{{0, 1}, {0, 4}, {1, 4}, {2, 3}}))
	assert.False(t, validTree(4, [][]int{{0, 1}, {2, 3}}))
	assert.True(t, validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))
	assert.False(t, validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}))
}

func TestFindCircleNum(t *testing.T) {
	assert.Equal(t, 1, findCircleNum([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
	assert.Equal(t, 2, findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
	assert.Equal(t, 3, findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
}

func TestMaxConsecutiveOnes(t *testing.T) {
	assert.Equal(t, 2, findMaxConsecutiveOnes([]int{1, 0}))
	//assert.Equal(t, 4, findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0}))
	//assert.Equal(t, 4, findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}

func TestValidSudoku(t *testing.T) {
	assert.False(t, isValidSudoku([][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'1', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '.', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))

	assert.True(t, isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '.', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))

	assert.False(t, isValidSudoku([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '6', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '8', '.', '.', '.', '.'},
		{'9', '.', '.', '.', '7', '5', '.', '.', '.'},
		{'.', '.', '.', '.', '5', '.', '.', '8', '.'},
		{'.', '.', '9', '.', '.', '.', '.', '.', '.'},
		{'2', '.', '6', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}))
}

func TestRotateArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(arr, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, arr)
}

func TestUniquePaths(t *testing.T) {
	require.Equal(t, 28, uniquePaths(3, 7))
	require.Equal(t, 3, uniquePaths(3, 2))
}

func TestTreeLevelOrder(t *testing.T) {
	tr := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.Equal(t, [][]int{{3}, {9, 20}, {7, 15}}, levelOrder(tr))
}

func TestRemoveNth(t *testing.T) {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	l := removeNthFromEnd(ln, 2)
	fmt.Println(l)
}

func TestInsertBST(t *testing.T) {
	bst := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}

	b := insertIntoBST(bst, 5)
	fmt.Println(b)
}

func TestLongestSubstring(t *testing.T) {
	assert.Equal(t, 2, lengthOfLongestSubstring("au"))
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbbbb"))
}

func TestInclusion(t *testing.T) {
	assert.True(t, checkInclusion("ab", "eidbaooo"))
	assert.False(t, checkInclusion("ab", "eidboaoo"))
	assert.True(t, checkInclusion("adc", "dcda"))
}

func TestValidBST(t *testing.T) {
	bst := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}

	assert.True(t, isValidBST(bst))
}

func TestSortColors(t *testing.T) {
	val := []int{2, 0, 2, 1, 1, 0}
	sortColors(val)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, val)
}

func TestMergeInterval(t *testing.T) {
	assert.Equal(t, [][]int{{1, 5}}, merge([][]int{{1, 4}, {4, 5}}))
	assert.Equal(t, [][]int{{0, 4}}, merge([][]int{{1, 4}, {0, 4}}))
	assert.Equal(t, [][]int{{1, 4}}, merge([][]int{{1, 4}, {2, 3}}))
}

func TestRotateMatrix(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(m)
	assert.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, m)
}

func TestGenerateMatrix(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, generateMatrix(3))
}

func TestPopulatingNextRight(t *testing.T) {
	n := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   5,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val:   6,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
	}

	connect(n)
}

func TestUpdateMatrix(t *testing.T) {
	m1 := [][]int{{0}, {0}, {0}, {0}}
	assert.Equal(t, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}, updateMatrix(m1))
}

func TestRottenFruits(t *testing.T) {
	assert.Equal(t, 4, orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}

func TestSearch2DMatrix(t *testing.T) {
	m := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	assert.True(t, searchMatrix(m, 5))
	assert.False(t, searchMatrix(m, 20))
}

func TestNonOverlappingIntervals(t *testing.T) {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	assert.Equal(t, 1, eraseOverlapIntervals(intervals))

	//intervals := [][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}}
	//assert.Equal(t, 3, eraseOverlapIntervals(intervals))
}

func TestTripletSub(t *testing.T) {
	assert.True(t, increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	assert.True(t, increasingTriplet([]int{1, 2, 3, 4, 5}))
	assert.True(t, increasingTriplet([]int{5, 4, 3, 2, 1}))
	assert.True(t, increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}

func TestProductArray(t *testing.T) {
	assert.Equal(t, []int{24, 12, 8, 6}, productExceptSelf([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{0, 0, 9, 0, 0}, productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func TestCombinations(t *testing.T) {
	res := combine(4, 2)
	_ = res
}

func TestPermutations(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute([]int{1, 2, 3}))
}
