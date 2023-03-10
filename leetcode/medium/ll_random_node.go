package medium

import (
	"math/rand"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	nums []int
}

func Constructor2(head *ListNode) Solution {
	next := head
	arr := make([]int, 0, 10)

	for next != nil {
		arr = append(arr, next.Val)
		next = next.Next
	}

	return Solution{
		nums: arr,
	}
}

func (s *Solution) GetRandom() int {
	num := rand.Int63n(int64(len(s.nums)))

	return s.nums[num]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
