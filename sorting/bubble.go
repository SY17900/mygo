package sorting

func BubbleSort(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		ifSwap := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				ifSwap = true
			}
		}
		if !ifSwap {
			return
		}
	}
}
