package backtracking

func IncreasingSSQ(nums []int) (ans [][]int) {
	var dfs func(i int)
	rec := make([]int, 0, len(nums))
	dfs = func(i int) {
		if len(rec) > 1 {
			temp := make([]int, len(rec))
			copy(temp, rec)
			ans = append(ans, temp)
		}
		used := map[int]struct{}{}
		for next := i; next < len(nums); next++ {
			if len(rec) != 0 && nums[next] < rec[len(rec)-1] {
				continue
			}
			if _, ok := used[nums[next]]; ok {
				continue
			}
			used[nums[next]] = struct{}{}
			rec = append(rec, nums[next])
			dfs(next + 1)
			rec = rec[:len(rec)-1]
		}
	}

	dfs(0)
	return
}
