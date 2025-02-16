package sorting

func down(nums []int, start int, end int) {
	parent := start
	child := 2*parent + 1
	for child <= end {
		if child+1 <= end && nums[child] < nums[child+1] {
			child++
		}
		if nums[parent] >= nums[child] {
			return
		} else {
			nums[parent], nums[child] = nums[child], nums[parent]
			parent, child = child, 2*child+1
		}
	}
}

func HeapSort(nums []int) {
	endIndex := len(nums) - 1
	for i := len(nums)/2 - 1; i >= 0; i-- {
		down(nums, i, endIndex)
	}
	for i := endIndex; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		down(nums, 0, i-1)
	}
}
