package medium

func ArrayOfProducts(array []int) []int {
	if len(array) == 2 {
		return array
	}
	res := make([]int, len(array))

	left := make([]int, len(array))
	right := make([]int, len(array))

	i := 1
	product := array[0]
	left[0] = 1
	for i < len(array) {
		left[i] = product
		product = product * array[i]
		i++
	}

	i = len(array) - 2
	product = array[len(array)-1]
	right[len(array)-1] = 1
	for i >= 0 {
		right[i] = product
		product = product * array[i]
		i--
	}

	for i := 0; i < len(array); i++ {
		res[i] = left[i] * right[i]
	}

	return res
}
