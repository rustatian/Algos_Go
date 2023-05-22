package medium

type Node22 struct {
	Next *Node22
	Val  int
}

type MyLinkedList struct {
	head   *Node22
	lenght int
}

func Constructor22() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if this.lenght == 0 || index < 0 || this.lenght <= index {
		return -1
	}

	cur := this.head
	curr := 0
	for cur.Next != nil {
		if curr == index {
			return cur.Val
		}

		curr++
		cur = cur.Next
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node22{
		Next: this.head,
		Val:  val,
	}

	this.head = node
	this.lenght++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.lenght == 0 || this.head == nil {
		this.AddAtHead(val)
		return
	}

	curr := this.head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &Node22{
		Val: val,
	}

	this.lenght++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.lenght || index == -1 {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	curr := this.head
	for i := 1; i < index; i++ {
		curr = curr.Next
	}

	node := &Node22{
		Next: curr.Next,
		Val:  val,
	}

	curr.Next = node
	this.lenght++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.lenght == 0 || this.lenght <= index {
		return
	}

	if index == 0 {
		this.head = this.head.Next
		this.lenght--
		return
	}

	curr := this.head

	for i := 1; i < index; i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next
	this.lenght--
}
