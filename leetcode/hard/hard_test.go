package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKSortedLists(t *testing.T) {
	l1 := &ListNode{
		Val: -2,
		Next: &ListNode{
			Val: -1,
			Next: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val:  -1,
					Next: nil,
				},
			},
		},
	}

	//l2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	//
	//l3 := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val:  6,
	//		Next: nil,
	//	},
	//}

	lists := make([]*ListNode, 2)
	lists[0] = l1
	lists[1] = nil
	//lists[2] = l3

	lres := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
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
		},
	}

	assert.Equal(t, lres, mergeKLists(lists))

	//res := make([]*ListNode, 1)
	//res[0] = nil
	//assert.Equal(t, nil, mergeKLists(res))
}

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
