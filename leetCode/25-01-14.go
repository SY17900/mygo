package leetcode

func MinOperationsEasy(nums []int, k int) (ans int) {
	for _, num := range nums {
		if num < k {
			ans++
		}
	}
	return
}
