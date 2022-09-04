package easy

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	mp := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			continue
		} else {
			mp = i
			break
		}
	}

	high := arr[:mp]
	low := arr[mp:]

	if len(high) == 0 || len(low) == 0 {
		return false
	}

	for i := 0; i < len(high)-1; i++ {
		if high[i] >= high[i+1] {
			return false
		}
	}

	for i := 0; i < len(low)-1; i++ {
		if low[i] <= low[i+1] {
			return false
		}
	}

	return true
}
