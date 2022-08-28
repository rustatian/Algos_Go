package easy

import (
	"math"
)

func findNumbers(nums []int) int {
	evenNums := 0
	for i := 0; i < len(nums); i++ {
		digits := math.Floor(math.Log10(float64(nums[i])) + 1)
		if int(digits)%2 == 0 {
			evenNums++
		}
	}
	return evenNums
}
