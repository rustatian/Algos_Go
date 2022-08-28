package easy

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	c := image[sr][sc]

	if color != c {
		dfs(image, sr, sc, c, color)
	}

	return image
}

func dfs(img [][]int, sr, sc, color, newColor int) {
	if img[sr][sc] == color {
		img[sr][sc] = newColor

		if sr >= 1 {
			dfs(img, sr-1, sc, color, newColor)
		}

		if sc >= 1 {
			dfs(img, sr, sc-1, color, newColor)
		}

		if sr+1 < len(img) {
			dfs(img, sr+1, sc, color, newColor)
		}
		if sc+1 < len(img[0]) {
			dfs(img, sr, sc+1, color, newColor)
		}

	}
}
