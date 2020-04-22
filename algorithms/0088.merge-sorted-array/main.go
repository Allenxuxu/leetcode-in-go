package algorithm

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	}
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	}
	i, j := len(nums1)-1, n-1
	for i >= 0 && j >= 0 && m-1 >= 0 {
		if nums1[m-1] > nums2[j] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[j]
			j--
		}

		i--
	}

	copy(nums1, nums2[:j+1])
}
