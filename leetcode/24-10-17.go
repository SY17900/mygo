package leetcode

import "slices"

// https://leetcode.cn/problems/count-the-number-of-inversions/description/

// dfs(i, j)表示i+1个数的所有排序中逆序对个数为j的排序个数
// 知道dfs(i-1, j)及以前的所有值
// 设perm[i]和perm[0]到perm[i-1]组成k个逆序对
// 枚举k = 0, 1, 2, ..., min(i, j): 每种选择对应dfs(i-1, j-k)的已知值

func NumberOfPermutations(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	f[0][0] = 1
	for i := 1; i < n; i++ {
		mx := m
		if req[i] >= 0 {
			mx = req[i]
		}
		if r := req[i-1]; r >= 0 {
			for j := r; j <= min(i+r, mx); j++ {
				f[i][j] = f[i-1][r]
			}
		} else {
			for j := 0; j <= mx; j++ {
				for k := 0; k <= min(i, j); k++ {
					f[i][j] = (f[i][j] + f[i-1][j-k]) % mod
				}
			}
		}
	}

	return f[n-1][req[n-1]]
}
