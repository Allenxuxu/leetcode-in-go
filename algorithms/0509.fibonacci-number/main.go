package algorithm

func fib(N int) int {
	dp := make([]int, N+3)

	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[N]
}

func fib1(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}

	return fib(N-1) + fib(N-2)
}
