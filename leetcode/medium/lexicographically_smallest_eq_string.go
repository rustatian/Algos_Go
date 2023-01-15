package medium

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	const MAX = 26
	root := [26]uint8{}

	for i := 0; i < MAX; i++ {
		root[i] = uint8(i)
	}

	n := len(s1)
	for i := 0; i < n; i++ {
		s1ch := s1[i] - 'a'
		s2ch := s2[i] - 'a'

		unionSm(s1ch, s2ch, &root)
	}

	res := ""

	for i := 0; i < len(baseStr); i++ {
		ch := findSm(baseStr[i]-'a', &root)
		res += string(ch + 'a')
	}

	return res
}

func findSm(x uint8, root *[26]uint8) uint8 {
	if x == root[x] {
		return x
	}

	root[x] = findSm(root[x], root)

	return root[x]
}

func unionSm(x, y uint8, root *[26]uint8) {
	rootX := findSm(x, root)
	rootY := findSm(y, root)

	if rootX == rootY {
		// union
		return
	}

	if rootX < rootY {
		root[rootY] = rootX
		return
	}

	root[rootX] = rootY
}
