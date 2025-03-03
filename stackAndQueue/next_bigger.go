package stackandqueue

func NextBigger0(nums1 []int, nums2 []int) []int {
	rec := make(map[int]int)
	stack := []int{}
	ans := make([]int, len(nums1))

	for _, num := range nums2 {
		for len(stack) != 0 && num > stack[len(stack)-1] {
			rec[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	for _, num := range stack {
		rec[num] = -1
	}

	for index, num := range nums1 {
		ans[index] = rec[num]
	}

	return ans
}

func NextBigger1(nums []int) []int {
	stack := []int{}
	ans := make([]int, len(nums))

	for index, num := range nums {
		for len(stack) != 0 && num > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}

	for _, num := range nums {
		for len(stack) != 0 && num > nums[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
	}

	for _, index := range stack {
		ans[index] = -1
	}

	return ans
}
