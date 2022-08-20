package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
