package leetcode

import "math"

func NumberOfPairs(nums1 []int, nums2 []int, k int) int {
	max1 := math.MinInt

	map1 := make(map[int]int)
	for _, num := range nums1 {
		map1[num]++
		max1 = max(max1, num)
	}

	map2 := make(map[int]int)
	for _, num := range nums2 {
		map2[num]++
	}

	var ans = 0

	for num, count2 := range map2 {
		for temp := num * k; temp <= max1; temp += num * k {
			if count1, exists := map1[temp]; exists {
				ans += count1 * count2
			}
		}
	}

	return ans
}
