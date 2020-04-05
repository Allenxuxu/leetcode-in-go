package algorithm

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
