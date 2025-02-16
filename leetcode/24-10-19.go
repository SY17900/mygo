package leetcode

func MinOperations(nums []int) int {
	ans := 0
	for _, num := range nums {
		if num == 0 && ans%2 == 0 {
			ans++
		}
		if num == 1 && ans%2 == 1 {
			ans++
		}
	}

	return ans
}
