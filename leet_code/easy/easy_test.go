package easy

import (
	"testing"

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

func TestMyQueue(t *testing.T) {
	q := Constructor()

	q.Push(1)
	q.Push(2)
	require.Equal(t, 1, q.Peek())
	require.Equal(t, 1, q.Pop())
	require.False(t, q.Empty())
}
