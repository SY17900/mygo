package leetcode

import "math"

func NumberOfPairs2(nums1 []int, nums2 []int, k int) (ans int64) {
	max1 := math.MaxInt

	map1 := make(map[int]int)
	for _, num := range nums1 {
		map1[num]++
		max1 = max(max1, num)
	}

	map2 := make(map[int]int)
	for _, num := range nums2 {
		map2[num]++
	}

	ans = 0

	for num2, count2 := range map2 {
		for temp := num2 * k; temp <= max1; temp += num2 * k {
			if count1, has := map1[temp]; has {
				ans += int64(count1) * int64(count2)
			}
		}
	}

	return
}
