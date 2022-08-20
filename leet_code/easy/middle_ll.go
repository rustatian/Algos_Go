package easy

func middleNode(head *ListNode) *ListNode {
	sp := head
	fp := head

	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}

	return sp
}
