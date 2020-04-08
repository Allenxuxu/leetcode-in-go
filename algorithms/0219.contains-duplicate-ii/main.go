package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 遍历数组，将 map 元素存入map，map 的 key 为元素值，value 为元素的索引下标
// 如果 map 中已经存在相同的 元素值，则计算其 下标差值，如果差值的绝对值 <= k ,return true
// 如果不满足，则覆盖 map 中原来的元素，因为旧值的索引值较小，再继续遍历，即便有第三个相同元素，其下标差值的绝对值也不会满足 <= k
func containsNearbyDuplicate(nums []int, k int) bool {
	tmp := make(map[int]*int, len(nums))

	for i, v := range nums {
		if tmp[v] != nil {
			diff := *tmp[v] - i
			if diff < 0 {
				diff = -diff
			}

			if diff <= k {
				return true
			} else {
				// 将 map 中原来的 value 覆盖
				tmpIndex := i
				tmp[v] = &tmpIndex
			}
		}

		// 不能直接 tmp[v] = &i , i 变量不是临时变量，在所有迭代中内存地址不变
		tmpIndex := i
		tmp[v] = &tmpIndex
	}

	return false
}
