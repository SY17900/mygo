package twopointer

import "sort"

func ThreeNumber(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			curSum := nums[i] + nums[left] + nums[right]
			if curSum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				for left < len(nums) && nums[left] == nums[left-1] {
					left++
				}
				right--
				for right >= 0 && nums[right] == nums[right+1] {
					right--
				}
			} else if curSum < target {
				left++
				for left < len(nums) && nums[left] == nums[left-1] {
					left++
				}
			} else {
				right--
				for right >= 0 && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return ans
}
