package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(ll *LinkedList) *LinkedList {
	node := ll
	for node != nil {
		next := node.Next
		for next != nil && node.Value == next.Value {
			next = next.Next
		}

		node.Next = next
		node = next
	}

	return ll
}
