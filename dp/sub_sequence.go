package dp

func LengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	st := []int{nums[0]}

	bs := func(t int) int {
		left, right := 0, len(st)-1
		for left < right {
			mid := left + (right-left)/2
			if t <= st[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		}

		return left
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > st[len(st)-1] {
			st = append(st, nums[i])
		} else if nums[i] < st[len(st)-1] {
			left := bs(nums[i])
			st[left] = nums[i]
		}
	}

	return len(st)
}

func LengthOfDuplicated(nums1, nums2 []int) int {
	ans := 0
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}

	return ans
}

func CommonSubsequence(text1, text2 string) int {
	ans := 0
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[j][i-1])
			}
			ans = max(ans, dp[i][j])
		}
	}

	return ans
}
