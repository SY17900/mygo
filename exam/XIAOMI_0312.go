package exam

import (
	"fmt"
	"sort"
)

func XIAOMI_EXAM1() {
	var n, k, ans int
	fmt.Scan(&n, &k)
	rec := make([]int, n)
	canDown := make([]int, 0, n-1)
	for i := range rec {
		fmt.Scan(&rec[i])
		if i > 0 && rec[i]-rec[i-1] > 2 {
			canDown = append(canDown, rec[i]-rec[i-1]-1)
		}
	}

	cnt := rec[n-1] - rec[0] + 1
	if cnt <= k {
		fmt.Println(ans + 2)
		return
	}

	sort.Ints(canDown)
	for i := len(canDown) - 1; i >= 0; i-- {
		ans += 2
		cnt -= canDown[i]
		if cnt <= k {
			fmt.Println(ans + 2)
			return
		}
	}
}

func XIAOMI_EXAM2() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)
	type car struct {
		people int
		thing  int
		cost   int
	}
	rec := []car{}

	for ; n > 0; n-- {
		var baseMoney, typeCount, baseX, baseY int
		fmt.Scan(&baseMoney, &typeCount, &baseX, &baseY)
		rec = append(rec, car{baseX, baseY, baseMoney})
		for range typeCount {
			var typeMoney, typeX, typeY int
			fmt.Scan(&typeMoney, &typeX, &typeY)
			rec = append(rec, car{baseX + typeX, baseY + typeY, baseMoney + typeMoney})
		}
	}

	dp := make([][]int, x+1)
	for i := range dp {
		dp[i] = make([]int, y+1)
	}
	cost, people, thing := rec[0].cost, rec[0].people, rec[0].thing
	for i := 1; i <= y; i++ {
		if i <= people {
			dp[i][0] = cost
		}
		if i > people {
			dp[i][0] = dp[i-people][0] + cost
		}
	}
	for j := 1; j <= x; j++ {
		if j <= thing {
			dp[0][j] = cost
		}
		if j > thing {
			dp[j][0] = dp[j-thing][0] + cost
		}
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if i-people < 0 && j-thing < 0 {
				dp[i][j] = cost
			} else if i-people < 0 {
				dp[i][j] = dp[0][j-thing] + cost
			} else if j-thing < 0 {
				dp[i][j] = dp[i-people][0] + cost
			} else {
				dp[i][j] = dp[i-people][j-thing] + cost
			}
		}
	}

	for car := 1; car < len(rec); car++ {
		cost, people, thing := rec[car].cost, rec[car].people, rec[car].thing
		for i := 0; i <= x; i++ {
			for j := 0; j <= y; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if i-people < 0 && j-thing < 0 {
					dp[i][j] = min(dp[i][j], cost)
				} else if i-people < 0 {
					dp[i][j] = min(dp[i][j], dp[0][j-thing]+cost)
				} else if j-thing < 0 {
					dp[i][j] = min(dp[i][j], dp[i-people][0]+cost)
				} else {
					dp[i][j] = min(dp[i][j], dp[i-people][j-thing]+cost)
				}
			}
		}
	}

	fmt.Println(dp[x][y])
}
