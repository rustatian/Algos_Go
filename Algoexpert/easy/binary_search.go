package main

func BinarySearch(array []int, target int) int {
	return binarySearchHelper(array, target, 0, len(array)-1)
}

func binarySearchHelper(array []int, target int, left, right int) int {
	middle := (left + right) / 2
	num := array[middle]
	if left > right {
		return -1
	}
	if num == target {
		return middle
	} else if num > target {
		return binarySearchHelper(array, target, left, middle - 1)
	} else {
		return binarySearchHelper(array, target, middle + 1, right)
	}
}
