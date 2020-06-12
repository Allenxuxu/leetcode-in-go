package algorithm

import (
	"sort"
)

// 排序
// 时间复杂度 O(n*log(n))
// 空间复杂度 O(1)
// 先排序，然后成对判断，步进2（i += 2, j += 2）
// 如果这一对不相等，则返回 nums[i]
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
// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 遍历数组，将 元素作为 v 存入标记， 如果发现当前 v 已经存在 则 delete 当前元素
// 遍历完成后， map 中只剩一个元素
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

// 异或操作 相同为 0 ，相异为 1
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func singleNumber2(nums []int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
