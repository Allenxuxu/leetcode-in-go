package algorithm

func splitArray(nums []int, m int) int {
	//  由于题目的返回要求：返回最小和的值
	//  最小和必然落在 [max(nums), sum(nums)] 之间
	//  我们可以使用二分来进行查找
	l, r := maxInArray(nums), sum(nums)
	for l < r {
		mid := l + (r-l)/2
		// 我们由前向后对nums进行划分，使其子数组和 <= mid，然后统计满足条件的数组数量
		// 若我们选的sum值过小，则满足条件的数量 > m，目标值应落在 [mid+1, high]
		// 若我们选的sum值过大，则满足条件的数量 < m，目标值应落在 [low, mid-1]
		var count, subSum int
		for i := 0; i < len(nums); i++ {
			subSum += nums[i]
			//  此时 subSum 不算 nums[i] 的话，刚好 <= mid
			if subSum > mid {
				count++
				subSum = nums[i]
			}
		}
		// 末尾还有一个子数组我们没有统计，这里把它加上
		count++
		if count > m {
			l = mid + 1
		} else {
			// 次数 count <= m, 我们使得 r = mid， 因为此时 mid 可能是答案，但也有可能会有更小的 mid 存在
			r = mid
		}
	}

	// 此时 l == r
	return r
}

func sum(nums []int) (ret int) {
	for _, v := range nums {
		ret += v
	}
	return
}

func maxInArray(nums []int) (max int) {
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return
}
