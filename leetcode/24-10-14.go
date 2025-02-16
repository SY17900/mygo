package leetcode

func SuperEggDrop(k int, n int) int {
	rec := make([][]int, n+1) //rec[i][j]：i层楼，j个鸡蛋对应的次数
	for i := range rec {
		rec[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		rec[i][1] = i
	}

	for i := 1; i <= n; i++ { // 枚举1-i，找到max(rec[x-1][j-1] + 1, rec[i-x][j] + 1)的最小值
		for j := 2; j <= k; j++ {
			l, r := 1, i
			for l < r {
				mid := (l + r + 1) >> 1
				a, b := rec[mid-1][j-1], rec[i-mid][j]
				if a <= b {
					l = mid
				} else {
					r = mid - 1
				}
			}
			rec[i][j] = max(rec[l-1][j-1], rec[i-l][j]) + 1
		}
	}

	return rec[n][k]
}
