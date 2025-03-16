package dp

func MaxProfit0(prices []int) int {
	dp := make([][]int, 2) // 0: hold, 1: unhold
	dp[0], dp[1] = make([]int, 2), make([]int, 2)
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], -prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], prices[i]+dp[(i-1)%2][0])
	}

	return dp[(len(prices)-1)%2][1]
}

func MaxProfit1(prices []int) int {
	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, 2), make([]int, 2)
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]-prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
	}

	return dp[(len(prices)-1)%2][1]
}

func MaxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, 5)
	dp[1] = -prices[0]
	dp[3] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[1] = max(dp[1], dp[0]-prices[i])
		dp[2] = max(dp[2], dp[1]+prices[i])
		dp[3] = max(dp[3], dp[2]-prices[i])
		dp[4] = max(dp[4], dp[3]+prices[i])
	}

	return dp[4]
}

func MaxProfit3(k int, prices []int) int {
	n := len(prices)
	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, k*2+1), make([]int, k*2+1)

	for i := 1; i <= k*2; i += 2 {
		dp[0][i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k*2; j++ {
			if j%2 == 1 {
				dp[i%2][j] = max(dp[(i-1)%2][j], dp[(i-1)%2][j-1]-prices[i])
			} else {
				dp[i%2][j] = max(dp[(i-1)%2][j], dp[(i-1)%2][j-1]+prices[i])
			}
		}
	}

	return dp[(n-1)%2][k*2]
}

func MaxProfit4(prices []int) int {
	state0, state1, state2, state3 := -prices[0], 0, 0, 0

	for i := 1; i < len(prices); i++ {
		temp0 := max(state0, max(state1, state3)+prices[i])
		temp1 := max(state1, state3)
		temp2, temp3 := state0+prices[i], state2

		state0, state1, state2, state3 = temp0, temp1, temp2, temp3
	}

	return max(state1, max(state2, state3))
}

func MaxProfit5(prices []int, fee int) int {
	dp := make([][]int, 2) // 0: hold; 1: unhold
	dp[0], dp[1] = make([]int, 2), make([]int, 2)
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i%2][0] = max(dp[(i-1)%2][0], dp[(i-1)%2][1]-prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i]-fee)
	}

	return dp[(len(prices)-1)%2][1]
}
