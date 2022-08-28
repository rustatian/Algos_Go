package easy

// saw

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		// save
		tmp := curr.Next

		// reverse
		curr.Next = prev

		// save prev
		prev = curr

		// restore
		curr = tmp
	}

	return prev
}
