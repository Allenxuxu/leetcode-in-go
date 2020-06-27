package algorithm

func searchRange(nums []int, target int) []int {
	ret := make([]int, 2)
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target { // 继续向左搜索，寻找左边界
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l >= len(nums) || nums[l] != target {
		l = -1
	}
	ret[0] = l

	l, r = 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target { // 继续向右边搜索，寻找右边界
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if r < 0 || nums[r] != target {
		r = -1
	}
	ret[1] = r

	return ret
}
