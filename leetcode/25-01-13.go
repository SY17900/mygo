package leetcode

func WaysToSplitArray(nums []int) (ans int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	start := 0
	for _, num := range nums[:len(nums)-1] {
		start += num
		if start*2 >= sum {
			ans++
		}
	}

	return
}
