package leetcode

import "math"

// 如果在a层没碎，在b层碎了，那么需要扔b-a-1次
// n层楼，假设在第m层扔了一次，如果碎了那就要再扔m-1次，否则扔1+f(m-n)
// 比如说4层：已知f(1)、f(2)和f(3)：f(4)=min(max(0, 1+f(3)), max(1, 1+f(2)), max(3, 1+f(3)))
func TowEggDrop(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	rec := make([]int, n+1)
	rec[0] = 0
	rec[1] = 1
	rec[2] = 2

	for i := 3; i <= n; i++ {
		ans := math.MaxInt
		for j := 1; j < i; j++ {
			ans = min(max(j, 1+rec[i-j]), ans)
		}
		rec[i] = ans
	}

	return rec[n]
}
