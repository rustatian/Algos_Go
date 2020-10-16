package main

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	multiplier := 1
	return productHelper(multiplier, array)
}

func productHelper(multiplier int, array []interface{}) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		// SpecialArray
		if sa, ok := array[i].(SpecialArray); ok {
			sum += productHelper(multiplier+1, sa)
		} else
		// number
		if n, ok := array[i].(int); ok {
			sum += n
		}
	}
	return sum * multiplier
}
