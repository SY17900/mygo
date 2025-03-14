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
