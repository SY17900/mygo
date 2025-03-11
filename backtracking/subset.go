package backtracking

import "sort"

func SubSet1(nums []int) (ans [][]int) {
	var dfs func(int)
	rec := make([]int, 0, len(nums))
	dfs = func(i int) {
		temp := make([]int, len(rec))
		copy(temp, rec)
		ans = append(ans, temp)
		for next := i; next < len(nums); next++ {
			rec = append(rec, nums[next])
			dfs(next + 1)
			rec = rec[:len(rec)-1]
		}
	}

	dfs(0)
	return
}

func SubSet2(nums []int) (ans [][]int) {
	var dfs func(int)
	rec := make([]int, 0, len(nums))
	dfs = func(i int) {
		temp := make([]int, len(rec))
		copy(temp, rec)
		ans = append(ans, temp)
		for next := i; next < len(nums); next++ {
			if next > i && nums[next] == nums[next-1] {
				continue
			}
			rec = append(rec, nums[next])
			dfs(next + 1)
			rec = rec[:len(rec)-1]
		}
	}

	sort.Ints(nums)
	dfs(0)
	return
}
