package algorithm

func myPow(x float64, n int) float64 {
	var ret float64
	var flag = 1
	if n <= 0 {
		flag = -1
		n = -n
	}
	if n == 0 {
		return 1
	}

	ret = myPow(x, n/2)

	if n%2 == 1 {
		ret = ret * ret * x
	} else {
		ret = ret * ret
	}

	if flag < 0 {
		ret = 1 / ret
	}
	return ret
}

func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var flag = 1
	if n < 0 {
		n = -n
		flag = -1
	}

	var ret float64 = 1
	for i := 0; i < n; i++ {
		ret *= x
	}

	if flag < 0 {
		ret = 1 / ret
	}
	return ret
}
