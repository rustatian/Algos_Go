package easy

//	[1,0,2,3,0,4,5,0]
//
// -> [1,0,0,2,3,0,0,4]
func duplicateZeros(arr []int) {
	if len(arr) == 1 {
		return
	}

	length := len(arr) - 1
	zeros := 0

	for i := 0; i <= length-zeros; i++ {
		if arr[i] == 0 {
			if i == length-zeros {
				arr[length] = 0
				length -= 1
				break
			}

			zeros++
		}
	}

	if zeros == 0 {
		return
	}

	last := length - zeros

	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+zeros] = 0
			zeros--
			arr[i+zeros] = 0
		} else {
			arr[i+zeros] = arr[i]
		}
	}
}
