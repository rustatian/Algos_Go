package main

import "fmt"

func shiftDown(numbers []int, k, n int) {
	// k << 1 is equal to k*2
	for k<<1 <= n {
		j := k << 1

		if j < n && numbers[j] < numbers[j+1] {
			j++
		}

		if !(numbers[k] < numbers[j]) {
			break
		}

		numbers[k], numbers[j] = numbers[j], numbers[k]
		k = j
	}
}

func heapSort(numbers []int) {
	arraySize := len(numbers) - 1

	for i := arraySize/2 - 1; i >= 0; i-- {
		shiftDown(numbers, i, arraySize)
	}

	for i := arraySize - 1; i >= 1; i-- {
		numbers[0], numbers[i] = numbers[i], numbers[0]
		shiftDown(numbers, 0, i-1)
	}

}

func main() {
	a := []int{2, 23, 33, 44, 1, 2, 2, 2, 4, 6, 99}

	heapSort(a)

	for _, m := range a {
		fmt.Println(m)
	}
}
