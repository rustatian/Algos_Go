package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
