package main

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)

	color := "B"
	// if red taller than blue, than red going to the front row
	if redShirtHeights[0] < blueShirtHeights[0] {
		color = "R"
	}

	for i := 0; i < len(redShirtHeights); i++ {
		if color == "R" {
			if redShirtHeights[i] >= blueShirtHeights[i] {
				return false
			}
		} else {
			if blueShirtHeights[i] >= redShirtHeights[i] {
				return false
			}
		}
	}

	return true
}
