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

func NumOfPalindrome(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				res++
				dp[i][j] = true
			}
		}
	}

	return res
}

func LongestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if i == j {
					dp[i][j] = 1
				} else if j-i == 1 {
					dp[i][j] = 2
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][len(s)-1]
}
