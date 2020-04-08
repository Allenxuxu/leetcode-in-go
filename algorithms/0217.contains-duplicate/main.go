package algorithm

import "sort"

// 时间复杂度 O(n * log n)
// 空间复杂度 O(1)
// 先排序，相同的元素就会排列在一起，再遍历识别
func containsDuplicate1(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 遍历数组，将元素插入 map ，如果已经存在，在发现重复
func containsDuplicate2(nums []int) bool {
	tmp := make(map[int]bool, len(nums))

	for k := range nums {
		if tmp[nums[k]] {
			return true
		}
		tmp[nums[k]] = true
	}

	return false
}
