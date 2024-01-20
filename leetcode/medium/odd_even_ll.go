package medium

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := &ListNode{
		Val:  0,
		Next: nil,
	}

	even := &ListNode{
		Val:  0,
		Next: nil,
	}

	curr := head
	odd_even := true

	for curr != nil {
		switch odd_even {
		case true:
			odd.Next = curr
			odd = odd.Next
		case false:
			even.Next = curr
			even = even.Next
		}

		curr = curr.Next
		odd_even = !odd_even
	}

	curr.Next = even

	return head
}
