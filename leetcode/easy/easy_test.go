package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, 999, longestPalindrome("abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa"))
	assert.Equal(t, 3, longestPalindrome("ccc"))
	assert.Equal(t, 7, longestPalindrome("abccccdd"))
	assert.Equal(t, 1, longestPalindrome("a"))
}

func TestTicTacToe(t *testing.T) {
	assert.Equal(t, "A", tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}))
}

func TestDesignOrderedStream(t *testing.T) {
	os := Constructor(5)
	assert.Equal(t, []string{}, os.Insert(3, "ccccc"))
	assert.Equal(t, []string{"aaaaa"}, os.Insert(1, "aaaaa"))
	assert.Equal(t, []string{"bbbbb", "ccccc"}, os.Insert(2, "bbbbb"))
	assert.Equal(t, []string{}, os.Insert(5, "eeeee"))
	assert.Equal(t, []string{"ddddd", "eeeee"}, os.Insert(4, "ddddd"))
}

func TestLargestPerimeter(t *testing.T) {
	assert.Equal(t, 5, largestPerimeter([]int{2, 1, 2}))
	assert.Equal(t, 0, largestPerimeter([]int{1, 2, 1}))
}

func TestContainsDuplicates(t *testing.T) {
	assert.True(t, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	assert.True(t, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	assert.False(t, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func TestIsValidPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("A man, a plan, a canal: Panama"))
	assert.True(t, isPalindrome("a."))
	assert.False(t, isPalindrome("0P"))
}

func TestValidPath(t *testing.T) {
	assert.True(t, validPath(10, [][]int{{2, 6}, {4, 7}, {1, 2}, {3, 5}, {7, 9}, {6, 4}, {9, 8}, {0, 1}, {3, 0}}, 3, 5))
	assert.True(t, validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	assert.False(t, validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
}

func TestFindPivotIDX(t *testing.T) {
	assert.Equal(t, 3, pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	assert.Equal(t, -1, pivotIndex([]int{1, 2, 3}))
	assert.Equal(t, 0, pivotIndex([]int{2, 1, -1}))
}

func TestSameTree(t *testing.T) {
	b1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	b2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	assert.True(t, isSameTree(b1, b2))
}

func TestThirdMax(t *testing.T) {
	assert.Equal(t, 2, thirdMax([]int{1, 2, 2, 5, 3, 5}))
	assert.Equal(t, 2, thirdMax([]int{1, 2}))
}

func TestFindDisappearedNumbers(t *testing.T) {
	assert.Equal(t, []int{7, 8}, findDisappearedNumbers([]int{3, 3, 2, 1, 4, 5, 6, 4}))
	assert.Equal(t, []int{5, 6}, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	assert.Equal(t, []int{2}, findDisappearedNumbers([]int{1, 1}))
}

func TestSortArrayByParity(t *testing.T) {
	assert.Equal(t, []int{2, 4, 3, 1}, sortArrayByParity([]int{3, 1, 2, 4}))
}

func TestMoveZeroes(t *testing.T) {
	in := []int{1, 0, 1, 0}
	moveZeroes(in)
	assert.Equal(t, []int{1, 1, 0, 0}, in)
	in = []int{0, 1, 0, 3, 12}
	moveZeroes(in)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, in)
}

func TestReplaceElementsWithGreatestElement(t *testing.T) {
	assert.Equal(t, []int{18, 6, 6, 6, 1, -1}, replaceElements([]int{17, 18, 5, 4, 6, 1}))
}

func TestValidMountainArray(t *testing.T) {
	assert.False(t, validMountainArray([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
	assert.False(t, validMountainArray([]int{3, 5, 5}))
	assert.True(t, validMountainArray([]int{0, 3, 2, 1}))
}

func TestCheckForDouble(t *testing.T) {
	assert.True(t, checkIfExist([]int{0, 0}))
	assert.False(t, checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
	assert.True(t, checkIfExist([]int{10, 2, 5, 3}))
	assert.True(t, checkIfExist([]int{7, 1, 14, 11}))
	assert.False(t, checkIfExist([]int{3, 1, 7, 11}))
}

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 2}))
	assert.Equal(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func TestRemoveElement(t *testing.T) {
	removeElement([]int{3, 3}, 3)
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	removeElement([]int{3, 2, 2, 3}, 3)
}

func TestMergeSortedArray(t *testing.T) {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{0}, 0, []int{1}, 1)
	merge([]int{1}, 1, nil, 0)
	merge([]int{2, 0}, 1, []int{1}, 1)
}

func TestDuplicateZeros(t *testing.T) {
	in := []int{8, 4, 5, 0, 0, 0, 0, 7}
	duplicateZeros(in)
	assert.Equal(t, []int{8, 4, 5, 0, 0, 0, 0, 0}, in)
}

func TestFindNumbers(t *testing.T) {
	assert.Equal(t, 1, findNumbers([]int{555, 901, 482, 1771}))
}

func TestConsecutiveOnes(t *testing.T) {
	assert.Equal(t, 2, findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}

func TestTwoSum(t *testing.T) {
	nums := make([]int, 0, 10)
	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 4)

	out := twoSum(nums, 6)
	require.Len(t, out, 2)
	require.Equal(t, 1, out[0])
	require.Equal(t, 2, out[1])
}

func TestTwoSum2(t *testing.T) {
	nums := make([]int, 0, 10)
	nums = append(nums, 3)
	nums = append(nums, 2)
	nums = append(nums, 3)

	out := twoSum(nums, 6)
	require.Len(t, out, 2)
	require.Equal(t, 2, out[0])
	require.Equal(t, 0, out[1])
}

func TestIntersect(t *testing.T) {
	require.Equal(t, []int{2, 2}, intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	require.Equal(t, []int{4, 9}, intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func TestReshapeMatrix(t *testing.T) {
	require.Equal(t, [][]int{{1}, {2}, {3}, {4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 4, 1))
	require.Equal(t, [][]int{{1, 2, 3, 4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	require.Equal(t, [][]int{{1, 2}, {3, 4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
}

func TestRansomNote(t *testing.T) {
	require.False(t, canConstruct("a", "b"))
	require.False(t, canConstruct("aa", "ab"))
	require.True(t, canConstruct("aa", "aab"))
}

func TestValidAnagram(t *testing.T) {
	require.True(t, isAnagram("anagram", "nagaram"))
	require.False(t, isAnagram("rat", "car"))
}

func TestHasCycle(t *testing.T) {
	ln := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: nil},
	}

	//ln.Next = &ListNode{Val: 2, Next: ln}

	hasCycle(ln)
}

func TestRemoveLLElement(t *testing.T) {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	removeElements(ln, 6)
}

func TestReverseLL(t *testing.T) {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	_ = reverseList(ln)
}

func TestRemoveDupLL(t *testing.T) {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	_ = deleteDuplicates(ln)
}

//func TestMyQueue(t *testing.T) {
//	q := Constructor()
//
//	q.Push(1)
//	q.Push(2)
//	require.Equal(t, 1, q.Peek())
//	require.Equal(t, 1, q.Pop())
//	require.False(t, q.Empty())
//}

func TestSquares(t *testing.T) {
	assert.Equal(t, []int{0, 1, 9, 16, 100}, sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func TestPreorder(t *testing.T) {
	tn := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	assert.Equal(t, []int{1, 2, 3}, preorderTraversal(tn))
}

func TestMaxDepth(t *testing.T) {
	//tn := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val:   9,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 20,
	//		Left: &TreeNode{
	//			Val: 3,
	//			Left: &TreeNode{
	//				Val:   15,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   7,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: nil,
	//	},
	//}
	//
	//assert.Equal(t, 3, maxDepth(tn))
	//
	//tn2 := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//
	//assert.Equal(t, 2, maxDepth(tn2))

	tn3 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	assert.Equal(t, 3, maxDepth(tn3))
}

func TestSymmetricTree(t *testing.T) {
	tn := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.True(t, isSymmetric(tn))

	tn2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.False(t, isSymmetric(tn2))

	tn3 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	assert.False(t, isSymmetric(tn3))
}

func TestPathSum(t *testing.T) {
	tn := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	assert.True(t, hasPathSum(tn, 22))

	tn2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	assert.False(t, hasPathSum(tn2, 1))
}

func TestInvertBT(t *testing.T) {
	tn := &TreeNode{
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
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	invertTree(tn)

	assert.Equal(t, &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}, tn)
}

func TestLLMiddle(t *testing.T) {
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
	l := middleNode(ln)
	fmt.Println(l)
}

func TestSearchBST(t *testing.T) {
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

	res := searchBST(bst, 5)
	fmt.Println(res)
}

func TestTwoSumBST(t *testing.T) {
	bst := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  6,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.True(t, findTarget(bst, 9))
	assert.False(t, findTarget(bst, 28))
}

func TestFloodFill(t *testing.T) {
	assert.Equal(t, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
}

func TestGetRow(t *testing.T) {
	assert.Equal(t, []int{1, 3, 3, 1}, getRow(3))
}

func TestMergeTwoBT(t *testing.T) {
	r1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	r2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	assert.Equal(t, &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}, mergeTrees(r1, r2))
}

func TestRunningSum(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 10}, runningSum([]int{1, 2, 3, 4}))
}
