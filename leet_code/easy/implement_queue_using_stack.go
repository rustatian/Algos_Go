package easy

// 6.55

type MyQueue struct {
	bs    []int
	start int
}

func ConstructorQ() MyQueue {
	return MyQueue{
		bs: make([]int, 0, 1),
	}
}

func (q *MyQueue) Push(x int) {
	q.bs = append(q.bs, x)
}

func (q *MyQueue) Pop() int {
	el := q.bs[q.start]
	q.start++
	return el
}

func (q *MyQueue) Peek() int {
	return q.bs[q.start]
}

func (q *MyQueue) Empty() bool {
	return q.start >= len(q.bs)
}
