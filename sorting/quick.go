package sorting

func partition(nums []int, left int, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot

	return left
}

func QuickSort(nums []int, left int, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		QuickSort(nums, left, pivot-1)
		QuickSort(nums, pivot+1, right)
	}
}
