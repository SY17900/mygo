package dp

func MaxProfit0(prices []int) int {
	dp := make([][]int, 2) // 0: hold, 1: unhold
	dp[0][0], dp[0][1] = -prices[0], 0

	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], -prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], prices[i]+dp[(i-1)%2][0])
	}

	return dp[(len(prices)-1)%2][1]
}

func MaxProfit1(prices []int) int {
	dp := make([][]int, 2)
	dp[0][0], dp[0][1] = -prices[0], 0

	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]-prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
	}

	return dp[(len(prices)-1)%2][1]
}
