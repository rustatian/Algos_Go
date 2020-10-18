package main

func BubbleSort(array []int) []int {
BEGIN:
	swapped := false
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			swapped = true
			tmp := array[i+1]
			array[i+1] = array[i]
			array[i] = tmp
		}
	}
	if swapped == true {
		goto BEGIN
	}

	return array
}
