package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{}
	tmp.Next = head

	num := 0

	count := head
	for count != nil {
		num++
		count = count.Next
	}

	counter := num - n
	count = tmp
	for counter > 0 {
		counter--
		count = count.Next
	}

	count.Next = count.Next.Next

	return tmp.Next
}
