package bw_97

import (
	"strconv"
)

func separateDigits(nums []int) []int {
	res := make([]int, 0, 10)

	for i := 0; i < len(nums); i++ {
		str := strconv.Itoa(nums[i])
		for j := 0; j < len(str); j++ {
			r, _ := strconv.Atoi(string(str[j]))
			res = append(res, r)
		}
	}

	return res
}
