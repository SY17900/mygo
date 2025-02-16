package slidingwindow

func FindMaxAverage(nums []int, k int) float64 {
	maxSum, windowSum := 0, 0
	for _, n := range nums[:k] {
		maxSum += n
		windowSum += n
	}

	for i := k; i < len(nums); i++ {
		windowSum -= nums[i-k] - nums[i]
		maxSum = max(windowSum, maxSum)
	}

	return float64(maxSum) / float64(k)
}
