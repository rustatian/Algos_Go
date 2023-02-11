package easy

type MyStack struct {
	q1 []int
}

func Constructor2() MyStack {
	return MyStack{
		q1: make([]int, 0, 10),
	}
}

func (ms *MyStack) Push(x int) {
	ms.q1 = append(ms.q1, x)
}

func (ms *MyStack) Pop() int {
	if len(ms.q1) == 0 {
		return 0
	}

	el := ms.q1[len(ms.q1)-1]
	ms.q1 = ms.q1[:len(ms.q1)-1]
	return el
}

func (ms *MyStack) Top() int {
	return ms.q1[len(ms.q1)-1]
}

func (ms *MyStack) Empty() bool {
	return len(ms.q1) == 0
}
