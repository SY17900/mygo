package dp

func BreakInteger(num int) int {
	dp := make([]int, num+1)
	dp[1] = 1
	for i := 2; i <= num; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[j]*(i-j), j*(i-j)))
		}
	}

	return dp[num]
}

func NumOfBST(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		temp := 0
		for j := 1; j <= i; j++ {
			temp += dp[j-1] * dp[i-j]
		}
		dp[i] = temp
	}

	return dp[n]
}
