package leetcode

import "slices"

// 计算candidates中的数字每种组合下按位与的结果，返回按位与结果大于0的最长组合的长度。

func LargestCombination(candidates []int) (ans int) {
	rec := [24]int{}
	for _, num := range candidates {
		for i := 0; num > 0; i++ {
			rec[i] += num & 1
			num >>= 1
		}
	}

	return slices.Max(rec[:])
}
