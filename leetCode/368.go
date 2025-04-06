package leetcode

import "slices"

func LargestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)
	n := len(nums)
	rec, from := make([]int, n), make([]int, n)
	for i := range from {
		from[i] = -1
	}
	maxI := 0

	for i, x := range nums {
		for j, y := range nums[:i] {
			if x%y == 0 && rec[j] > rec[i] {
				rec[i] = rec[j]
				from[i] = j
			}
		}
		rec[i]++
		if rec[i] > rec[maxI] {
			maxI = i
		}
	}

	ans := make([]int, 0, rec[maxI])
	for i := maxI; i >= 0; i = from[i] {
		ans = append(ans, nums[i])
	}
	return ans
}
