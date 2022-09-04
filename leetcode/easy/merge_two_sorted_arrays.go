package easy

/*
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	pWr := m + n - 1
	p1 := m - 1
	p2 := n - 1

	for p2 >= 0 {
		if pWr < 0 {
			break
		}

		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[pWr] = nums1[p1]
			p1--
		} else {
			nums1[pWr] = nums2[p2]
			p2--
		}

		pWr--
	}

	return
}
