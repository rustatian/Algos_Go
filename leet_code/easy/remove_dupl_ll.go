package easy

// 12.33 - see

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}

	}

	return head
}
