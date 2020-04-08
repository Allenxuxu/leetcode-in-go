package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// 从地位到高位逐个判断进位的情况
// 需要特别注意的情况，是最高如果也需要进位置，则需要扩容数组并拷贝，最高位 为 1
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 > 9 {
			digits[i] = 0
			if i == 0 {
				digits = append(digits[:1], digits...)
				digits[0] = 1
				// 以下写法作用相同
				//ret := make([]int, len(digits)+1)
				//copy(ret[1:], digits)
				//ret[0] = 1
				//return ret
			}
		} else {
			digits[i] += 1
			break
		}
	}
	return digits
}
