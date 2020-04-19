package algorithm

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	var ret [][]int
	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}
	for k := 0; k <= len(nums)-3; k++ {
		for i, j := k+1, len(nums)-1; i < j; {
			ab := nums[i] + nums[j]
			if -nums[k] == ab {
				ret = append(ret, []int{nums[k], nums[i], nums[j]})
				for ; i < j && nums[i+1] == nums[i]; i++ {
				}
				for ; i < j && nums[j-1] == nums[j]; j-- {

				}
				i++
				j--
			} else if -nums[k] < nums[i]+nums[j] {
				j--
			} else {
				i++
			}
		}

		for ; k < len(nums)-1 && nums[k+1] == nums[k]; k++ {
		}
	}

	return ret
}

func threeSum1(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	ret := make([][]int, 0)
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			b, c := nums[i], nums[j]
			a := -c - b
			v, ok := m[a]
			if ok {
				if a == b && b == c && m[a] < 3 {
					continue
				}
				if b == c && m[b] < 2 {
					continue
				}
				if (a == b || a == c) && v < 2 {
					continue
				}

				retItem := []int{a, b, c}
				sort.Ints(retItem)
				for _, v := range ret {
					if repeat(v, retItem) {
						goto exit
					}
				}
				ret = append(ret, retItem)
			}
		exit:
		}
	}

	return ret
}

// 保证 a, b 排序过，且len 相等
func repeat(a []int, b []int) bool {
	var ret = true
	for k, v := range a {
		if v != b[k] {
			ret = false
			break
		}
	}
	return ret
}
