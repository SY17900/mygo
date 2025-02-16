package leetcode

import (
	"math"
	"sort"
)

func MinimumAverage(nums []int) float64 {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	ans := math.MaxInt
	for i < j {
		ans = min(ans, nums[i]+nums[j])
		i++
		j--
	}

	return float64(ans) / 2
}
