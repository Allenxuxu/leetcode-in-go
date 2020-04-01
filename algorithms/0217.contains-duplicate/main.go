package algorithm

import "sort"

func containsDuplicate1(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	tmp := make(map[int]bool, len(nums))

	for k, _ := range nums {
		if tmp[nums[k]] {
			return true
		}
		tmp[nums[k]] = true
	}

	return false
}
