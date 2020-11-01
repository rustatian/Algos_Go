package medium

func IsMonotonic(array []int) bool {
	if len(array) == 0 || len(array) == 1 {
		return true
	}

	decreasing := true
	increasing := true

	for i := 0; i < len(array)-1; i++ {
		if array[i] < array[i+1] {
			decreasing = false
		}

		if array[i] > array[i+1] {
			increasing = false
		}
	}
	// Logical OR
	// TRUE if one of the nonDecreasing or nonIncreasing is TRUE
	// FALSE if both are false
	// WHY?
	// If we would have mixed increasing and decreasing values in the array, both decreasing and increasing will be false
	// and we will have false || false. If array is monolithic we will have false only in decreasing OR increasing. And the result
	// will be TRUE || FALSE --> TRUE
	return decreasing || increasing
}
