package dp

import (
	"math"
)

func Backpack01(costs, values []int, vol int) int {
	dp := make([]int, vol+1)
	for i := range costs {
		for j := vol; j >= costs[i]; j-- {
			dp[j] = max(dp[j], dp[j-costs[i]]+values[i])
		}
	}

	return dp[vol]
}

func DivideSet(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2

	dp := make([]int, sum+1)
	for _, num := range nums {
		for j := sum; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
		if dp[sum] == sum {
			return true
		}
	}

	return false
}

func MaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for _, str := range strs {
		num0, num1 := 0, 0
		for _, chr := range str {
			if chr == '0' {
				num0++
			} else {
				num1++
			}
		}

		for i := m; i >= num0; i-- {
			for j := n; j >= num1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-num0][j-num1]+1)
			}
		}
	}
	return dp[m][n]
}

func Change0(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func Change1(amount int, coins []int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func CombinationSum(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

func WordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		set[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := range i {
			if dp[j] && set[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
