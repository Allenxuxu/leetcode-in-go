package algorithm

// 要爬到第 n 阶，有两种方式：
// 从第 n - 1 阶上去(一次跨 1 阶)
// 从第 n - 2 阶上去(一次跨 2 阶)
// 所以 climbStairs(n) = climbStairs(n-1) + climbStairs(n-2)
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

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs1(n int) int {
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

func climbStairs2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	rec := make([]int, n+1)
	rec[1], rec[2] = 1, 2

	for i := 3; i <= n; i++ {
		rec[i] = rec[i-1] + rec[i-2]
	}

	return rec[n]
}
