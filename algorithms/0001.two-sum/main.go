package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 将数组中所有元素插入 map 中， key 为数组元素的值， value 为数组下标
// 遍历数组，用 target 减去 v， 如果其差值存在于 map 中则记录两者 index 退出循环
func twoSum(nums []int, target int) []int {
	var ret []int
	m := make(map[int]interface{}, len(nums))
	for k, v := range nums {
		m[v] = k
	}

	for k, v := range nums {
		need := target - v
		if m[need] != nil && m[need].(int) != k {
			ret = append(ret, k, m[need].(int))
			break
		}
	}

	return ret
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 解法同上，将两次遍历合并
func twoSum1(nums []int, target int) []int {
	var ret []int
	m := make(map[int]interface{}, len(nums))
	for k, v := range nums {
		need := target - v
		if m[need] != nil && m[need].(int) != k {
			ret = append(ret, m[need].(int), k)
			break
		}

		m[v] = k
	}

	return ret
}
