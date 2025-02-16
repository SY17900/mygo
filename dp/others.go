package dp

func CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	rec := make([]bool, target+1) // 0
	rec[0] = true
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			if rec[j-nums[i]] {
				rec[j] = true
			}
		}
	}

	return rec[target]
}
