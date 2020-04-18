package algorithm

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	var f1, f2 = 1, 2
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}
