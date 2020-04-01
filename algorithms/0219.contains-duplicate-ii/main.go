package algorithm

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
