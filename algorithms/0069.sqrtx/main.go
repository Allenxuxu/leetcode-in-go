package algorithm

func mySqrt(x int) int {
	var ret int
	s, e := 0, x
	for s <= e {
		ret = (s + e) / 2
		tmp := ret * ret
		if tmp <= x {
			if (ret+1)*(ret+1) > x {
				break
			} else {
				s = ret + 1
			}
		} else {
			e = ret - 1
		}
	}

	return ret
}
