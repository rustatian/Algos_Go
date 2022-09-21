package medium

import (
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	root := make([]int, len(s))
	rank := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		root[i] = i
		rank[i] = 1
	}

	resStr := make([]byte, len(s))

	// union all
	for i := 0; i < len(pairs); i++ {
		union5(pairs[i][0], pairs[i][1], root, rank)
	}

	hm := make(map[int]*[]int)

	/*
		here we need to create a sort-groups, for example:
		letters within indexes [0,2,7,9] are connected via graph
		then we need to sort them and put back to the original place but sorted
	*/
	for i := 0; i < len(root); i++ {
		// idx
		if v, ok := hm[find5(root[i], root)]; ok {
			*v = append(*v, i)
			continue
		}

		hm[find5(root[i], root)] = &[]int{i}
	}

	// if we have only 1 group, all letters are connected to each other
	// so, just sort the whole string and return the answer
	if len(hm) == 1 {
		r := []byte(s)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		return string(r)
	}

	// go over all index groups
	for _, v := range hm {
		// if only 1 letter -> just put it where it should be in the original string
		if len(*v) == 1 {
			resStr[(*v)[0]] = s[(*v)[0]]
			continue
		}

		// []byte used here as an optimization
		var str []byte

		// copy
		for i := 0; i < len(*v); i++ {
			str = append(str, s[(*v)[i]])
		}

		// sort []byte str
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})

		// put back, but sorted
		for i := 0; i < len(*v); i++ {
			resStr[(*v)[i]] = str[i]
		}
	}

	// return "sorted" string
	return string(resStr)
}

func find5(x int, root []int) int {
	if root[x] == x {
		return x
	}

	root[x] = find5(root[x], root)

	return root[x]
}

func union5(x, y int, root []int, rank []int) {
	rootX := find5(x, root)
	rootY := find5(y, root)

	if rootX == rootY {
		return
	}

	if rank[rootX] > rank[rootY] {
		root[rootY] = rootX
	} else if rank[rootX] < rank[rootY] {
		root[rootX] = rootY
	} else {
		root[rootY] = rootX
		rank[rootX]++
	}

	return
}
