package easy

// https://en.wikipedia.org/wiki/Sentinel_node
func removeElements(head *ListNode, val int) *ListNode {
	ll := &ListNode{
		Val:  0,
		Next: nil,
	}

	ll.Next = head

	// save prev
	prev, curr := ll, head

	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}

		curr = curr.Next
	}

	// skip first 0
	return ll.Next
}
