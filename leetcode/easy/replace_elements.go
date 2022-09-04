package easy

func replaceElements(arr []int) []int {
	ret := make([]int, len(arr))
	ret[len(ret)-1] = -1

	max := arr[len(arr)-1]

	for i := len(arr) - 2; i >= 0; i-- {
		ret[i] = max

		if arr[i] > max {
			max = arr[i]
		}
	}

	return ret
}
