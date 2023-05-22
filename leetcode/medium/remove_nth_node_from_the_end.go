package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	node := head
	length := 0

	for node != nil {
		node = node.Next
		length++
	}

	if length == n {
		return head.Next
	}

	target := length - n
	node = head

	for target != 0 {
		node = node.Next
		target--
	}

	node.Next = node.Next.Next

	return node
}
