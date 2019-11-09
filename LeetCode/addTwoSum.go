package main

//You are given two non-empty linked lists representing two non-negative integers.
//	The digits are stored in reverse order and each of their nodes contain a single digit.
//	Add the two numbers and return it as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Example:
//
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sl1 := make([]int, 10)
	sl2 := make([]int, 10)


	for {
		if l1.Next == nil {

		}



		break
	}


	return ln
}
