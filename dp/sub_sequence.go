package dp

func LengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	st := []int{nums[0]}

	bs := func(t int) int {
		left, right := 0, len(st)-1
		for left < right {
			mid := left + (right-left)/2
			if t <= st[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		}

		return left
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > st[len(st)-1] {
			st = append(st, nums[i])
		} else if nums[i] < st[len(st)-1] {
			left := bs(nums[i])
			st[left] = nums[i]
		}
	}

	return len(st)
}
