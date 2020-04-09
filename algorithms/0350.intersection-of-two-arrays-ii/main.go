package algorithm

import "sort"

// 时间复杂度 O(n log n)
// 空间复杂度 O(1)
// 先将两个数组排序，然后使用双指针遍历找到 nums1[i] == nums2[j] 时将 nums1[k] = nums1[i]
// 最后返回 nums1[k]
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

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// map 的 key 为数组元素的值，value 为数组元素出现的次数
// 找出较小的一个数组，将其元素插入到 map 中
// 遍历较大的数组，如果其元素在 map 中出现，则将 map 的value 减1，并将元素插入准备返回的数组
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
