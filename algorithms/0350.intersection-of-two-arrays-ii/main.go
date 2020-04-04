package algorithm

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var k int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			k++
			i++
			j++
		}
	}
	return nums1[:k]
}

func intersect1(nums1 []int, nums2 []int) []int {
	var (
		ret    []int
		length int
		short  []int
		long   []int
	)
	if len(nums1) < len(nums2) {
		length = len(nums1)
		short = nums1
		long = nums2
	} else {
		length = len(nums2)
		short = nums2
		long = nums1
	}

	tmp := make(map[int]int, length)

	for _, v := range short {
		tmp[v] = tmp[v] + 1
	}
	for _, v := range long {
		if tmp[v] != 0 {
			tmp[v] = tmp[v] - 1
			ret = append(ret, v)
		}
	}

	return ret
}
