package stack

func NextHigher(nums []int) []int {
	ans := make([]int, len(nums))
	stack := []int{}

	for index, num := range nums {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < num {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[top] = index - top
		}
		stack = append(stack, num)
	}

	return ans
}
