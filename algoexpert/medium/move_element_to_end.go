package medium

// 2, 1, 2, 2, 2, 3, 4, 2 --> 1, 3, 4, 2, 2, 2, 2, 2
//
func MoveElementToEnd(array []int, toMove int) []int {
	right := len(array) - 1
	left := 0

	for left <= right {
		// right element is equal toMove, continue search
		if array[right] == toMove {
			right--
			continue
		}
		// left element is not toMove, searching when right will not be toMove, and left will be toMove to swap
		if array[left] != toMove {
			left++
			continue
		}

		// at this point we know, that array[right] is not toMove and array[left] is toMove and we swap it and increment both indexes
		tmp := array[right]
		array[right] = toMove
		array[left] = tmp
		right--
		left++
	}
	return array
}
