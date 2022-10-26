package hard

import (
	"math"
	"sort"
)

func largestComponentSize(nums []int) int {
	sort.Ints(nums)

	root := make([]int, nums[len(nums)-1]+1)
	rank := make([]int, nums[len(nums)-1]+1)

	for i := 0; i < nums[len(nums)-1]+1; i++ {
		root[i] = i
		rank[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for factor := 2; factor < int(math.Sqrt(float64(nums[i]))+1); factor++ {
			if nums[i]%factor == 0 {
				union(nums[i], factor, root, rank)
				union(nums[i], nums[i]/factor, root, rank)
			}
		}
	}

	size := 0
	hm := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := find(nums[i], root)
		count := hm[num]
		count++
		hm[num] = count
		size = max(size, count)
	}

	return size
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func find(x int, root []int) int {
	if x == root[x] {
		return x
	}

	root[x] = find(root[x], root)

	return root[x]
}

func union(x, y int, root []int, rank []int) int {
	rootX := find(x, root)
	rootY := find(y, root)

	if rootX == rootY {
		// union
		return rootX
	}

	if rank[rootX] > rank[rootY] {
		t := rootX
		rootX = rootY
		rootY = t
	}

	// connect
	root[rootX] = rootY
	rank[rootY] += rank[rootX]

	return rootY
}
