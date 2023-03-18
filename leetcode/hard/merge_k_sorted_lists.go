package hard

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodes []*ListNode

func (p *ListNodes) Len() int {
	return len(*p)
}

func (p *ListNodes) Less(i, j int) bool {
	return (*p)[i].Val < (*p)[j].Val
}

func (p *ListNodes) Pop() interface{} {
	val := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return val
}

func (p *ListNodes) Push(item interface{}) {
	*p = append(*p, item.(*ListNode))
}

func (p *ListNodes) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 || lists == nil {
		return nil
	}

	p := make(ListNodes, 0, 100)

	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}

		tmp := lists[i]
		for tmp != nil {
			p.Push(tmp)
			tmp = tmp.Next
		}
	}

	heap.Init(&p)

	if p.Len() == 0 {
		return nil
	}

	node := heap.Pop(&p).(*ListNode)
	head := &ListNode{
		Val:  node.Val,
		Next: nil,
	}

	tail := head

	for p.Len() > 0 {
		node := heap.Pop(&p).(*ListNode)
		tail.Next = node
		tail = node
	}

	tail.Next = nil

	return head
}
