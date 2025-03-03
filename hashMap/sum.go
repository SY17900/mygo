package hashmap

func TwoNumber(nums []int, target int) []int {
	rec := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := rec[target-num]; ok {
			return []int{index, i}
		}
		rec[num] = i
	}

	return []int{}
}

func FourArraysSum(nums1, nums2, nums3, nums4 []int) (ans int) {
	rec := make(map[int]int)

	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			rec[num1+num2]++
		}
	}

	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			if cnt, ok := rec[0-num3-num4]; ok {
				ans += cnt
			}
		}
	}

	return
}
