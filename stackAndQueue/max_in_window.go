package stackandqueue

func MaxSlidingWindow(nums []int, k int) (ans []int) {
	q := make([]int, 0, k)
	for i, num := range nums {
		for len(q) > 0 && num >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-k >= q[0] {
			q = q[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}

	return
}
