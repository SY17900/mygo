package backtracking

func CombineK(n, k int) (ans [][]int) {
	var dfs func(int)
	rec := make([]int, 0, k)
	dfs = func(i int) {
		if len(rec) == k {
			temp := make([]int, k)
			copy(temp, rec)
			ans = append(ans, temp)
			return
		}
		for next := i; next <= n; next++ {
			rec = append(rec, next)
			dfs(next + 1)
			rec = rec[:len(rec)-1]
		}
	}

	dfs(1)
	return
}
