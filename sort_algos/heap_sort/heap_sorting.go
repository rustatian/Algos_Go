package main

import "fmt"

func siftDown(numbers []int, root, bottom int) {
	var maxChild int
	var done bool

	// root << 1 equals to root * 2
	for (root<<1 <= bottom) && !done {
		if root<<1 == bottom {
			maxChild = root << 1
		} else if numbers[root<<1] > numbers[(root<<1)|1] {
			maxChild = root << 1
		} else {
			maxChild = (root << 1) | 1
		}

		if numbers[root] < numbers[maxChild] {
			numbers[root], numbers[maxChild] = numbers[maxChild], numbers[root]
			root = maxChild
		} else {
			done = true
		}
	}
}

func heapSort(numbers []int) {
	arraySize := len(numbers) - 1

	for i := arraySize/2 - 1; i >= 0; i-- {
		siftDown(numbers, i, arraySize)
	}

	for i := arraySize - 1; i >= 1; i-- {
		numbers[0], numbers[i] = numbers[i], numbers[0]
		siftDown(numbers, 0, i-1)
	}

}

func main() {
	a := []int{2, 23, 33, 44, 1, 2, 2, 2, 4, 6, 99}

	heapSort(a)

	for _, m := range a {
		fmt.Println(m)
	}
}
