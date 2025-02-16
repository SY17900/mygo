package leetcode

func DuplicateNumbersXOR(nums []int) (ans int) {
	var rec int = 0

	for _, num := range nums {
		if rec&(1<<num) != 0 {
			ans ^= num
		} else {
			rec |= (1 << num)
		}
	}

	return
}
