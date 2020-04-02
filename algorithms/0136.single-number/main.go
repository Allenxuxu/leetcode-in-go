package algorithm

import (
	"sort"
)

// 排序
// 时间复杂度 不为 O(n)
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)
	for i, j := 0, 1; j < len(nums); i, j = i+2, j+2 {
		if nums[i] != nums[j] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}

// 数学方法
// 2∗(a+b+c)−(a+a+b+b+c)=c
// 实现： 略

// map
// 空间复杂度 不为 O(1)
func singleNumber1(nums []int) int {
	tmp := make(map[int]interface{})
	for _, v := range nums {
		if tmp[v] != nil {
			delete(tmp, v)
		} else {
			tmp[v] = v
		}
	}

	// 此时 map 中只有一个元素
	for k := range tmp {
		return k
	}

	// 不会执行到这里
	return -1
}

// 异或操作
func singleNumber2(nums []int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
