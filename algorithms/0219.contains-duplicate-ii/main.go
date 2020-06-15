package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 遍历数组，将 map 元素存入map，map 的 key 为元素值，value 为元素的索引下标
// 如果 map 中已经存在相同的 元素值，则计算其 下标差值，如果差值的绝对值 <= k ,return true
// 如果不满足，则覆盖 map 中原来的元素，因为旧值的索引值较小，再继续遍历，即便有第三个相同元素，其下标差值的绝对值也不会满足 <= k
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))

	for index, value := range nums {
		i, ok := m[value]
		if ok {
			if index-i <= k {
				return true
			}
		}

		m[value] = index
	}

	return false
}
