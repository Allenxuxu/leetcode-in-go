package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 将数组中所有元素插入 map 中， key 为数组元素的值， value 为数组下标
// 遍历数组，用 target 减去 v， 如果其差值存在于 map 中则记录两者 index 退出循环
// 注意：如果存在重复元素，构建 map 时，数组下表较大的会覆盖较小的。而第二次遍历时，会先取出较小下表的元素，因此可以通过case。
func twoSum(nums []int, target int) []int {
	var ret []int
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		m[v] = k
	}

	for k, v := range nums {
		need := target - v
		if e, ok := m[need]; ok && e != k {
			ret = append(ret, k, e)
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
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		// 先比较，在赋值，避免重复元素被覆盖
		need := target - v
		if e, ok := m[need]; ok && e != k {
			ret = append(ret, e, k)
			break
		}

		m[v] = k
	}

	return ret
}
