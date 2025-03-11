package backtracking

import "sort"

func FullPermutation(nums []int) (ans [][]int) {
	rec := make([]int, 0, len(nums))
	var dfs func()
	used := make([]bool, len(nums))
	dfs = func() {
		if len(rec) == len(nums) {
			temp := make([]int, len(rec))
			copy(temp, rec)
			ans = append(ans, temp)
			return
		}
		for next := range nums {
			if used[next] {
				continue
			}
			used[next] = true
			rec = append(rec, nums[next])
			dfs()
			rec = rec[:len(rec)-1]
			used[next] = false

		}
	}

	dfs()
	return
}

func FullPermutationD1(nums []int) (ans [][]int) {
	rec := make([]int, 0, len(nums))
	var dfs func()
	used := make([]bool, len(nums))
	dfs = func() {
		if len(rec) == len(nums) {
			temp := make([]int, len(rec))
			copy(temp, rec)
			ans = append(ans, temp)
			return
		}
		levelUsed := map[int]struct{}{}
		for next := range nums {
			if used[next] {
				continue
			}
			if _, ok := levelUsed[nums[next]]; ok {
				continue
			}
			levelUsed[nums[next]] = struct{}{}
			used[next] = true
			rec = append(rec, nums[next])
			dfs()
			rec = rec[:len(rec)-1]
			used[next] = false
		}
	}

	dfs()
	return
}

func FullPermutationD2(nums []int) (ans [][]int) {
	var dfs func()
	rec := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	dfs = func() {
		if len(rec) == len(nums) {
			temp := make([]int, len(rec))
			copy(temp, rec)
			ans = append(ans, temp)
			return
		}
		for next := range nums {
			if next > 0 && nums[next] == nums[next-1] && !used[next-1] {
				continue
			}
			used[next] = true
			rec = append(rec, nums[next])
			dfs()
			rec = rec[:len(rec)-1]
			used[next] = false
		}
	}

	sort.Ints(nums)
	dfs()
	return
}
