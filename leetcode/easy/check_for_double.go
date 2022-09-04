package easy

func checkIfExist(arr []int) bool {
	hm := make(map[int]struct{}, len(arr))

	for i := 0; i < len(arr); i++ {
		if _, ok := hm[arr[i]*2]; ok {
			return true
		}

		// numbers like 5,7
		if arr[i]%2 != 0 {
			hm[arr[i]] = struct{}{}
			continue
		}

		// 10, 12 etc
		if _, ok := hm[arr[i]/2]; ok {
			return true
		}

		hm[arr[i]] = struct{}{}
	}

	return false
}
