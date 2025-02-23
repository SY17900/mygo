package stack

func NextHigher(nums []int) []int {
	ans := make([]int, len(nums))
	rec := []int{}

	for index, num := range nums {
		if len(rec) != 0 && nums[rec[len(rec)-1]] < num {
			top := rec[len(rec)-1]
			rec = rec[:len(rec)-1]

			ans[top] = index - top
		}
	}

	return ans
}
