package stackandqueue

func LargestRectangleArea(heights []int) int {
	stack := []int{0}
	ans := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	for i := 1; i < len(heights); i++ {

		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				w := i - stack[len(stack)-1] - 1
				h := heights[mid]
				ans = max(ans, h*w)
			}
		}

		if heights[i] == heights[stack[len(stack)-1]] {
			stack[len(stack)-1] = i
		} else {
			stack = append(stack, i)
		}
	}

	return ans
}
