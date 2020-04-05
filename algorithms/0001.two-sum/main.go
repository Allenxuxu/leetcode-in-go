package algorithm

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
