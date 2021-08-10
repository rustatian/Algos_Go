package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	ll := linkedList

	for ll != nil {
		if ll.Next != nil {

			if ll.Value == ll.Next.Value {
				ll = ll.Next
			}

			ll = ll.Next
		}
	}

	return ll
}
