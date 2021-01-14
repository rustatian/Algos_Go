package main

func BinarySearch(array []int, target int) int {
	return binarySearchHelperIterative(array, target, 0, len(array)-1)
}

func binarySearchHelper(array []int, target int, left, right int) int {
	if left > right {
		return -1
	}
	middle := (left + right) / 2
	num := array[middle]

	if num == target {
		return middle
	} else if num > target {
		return binarySearchHelper(array, target, left, middle-1)
	} else {
		return binarySearchHelper(array, target, middle+1, right)
	}
}

func binarySearchHelperIterative(array []int, target int, left, right int) int {
	res := -1
	for left <= right {
		middle := (left + right) / 2
		num := array[middle]
		if num == target {
			return middle
		} else if num > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return res
}
