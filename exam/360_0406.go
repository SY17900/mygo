package exam

import "fmt"

func SANLIULING_EXAM1() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	var helper func(int, int) int
	helper = func(start, parts int) int {
		if parts == 1 {
			sum := 0
			mp := make(map[int]struct{})
			for _, num := range nums[start:] {
				if _, ok := mp[num]; !ok {
					sum += num
					mp[num] = struct{}{}
				}
			}
			return sum
		}

		sum := 0
		for i := start; i <= len(nums)-parts; i++ {
			tempSum := 0
			mp := make(map[int]struct{})
			for _, num := range nums[start : i+1] {
				if _, ok := mp[num]; !ok {
					tempSum += num
					mp[num] = struct{}{}
				}
			}
			tempSum += helper(i+1, parts-1)
			sum = max(sum, tempSum)
		}

		return sum
	}

	helper(0, k)
}

func SANLIULING_EXAM2() {
	var n, m int
	fmt.Scan(&n, &m)
	buys := make([]int, n)
	for i := range buys {
		fmt.Scan(&buys[i])
	}

	has := 0
	for range m {
		buyNum := has % n
		has += buys[buyNum]
	}

	fmt.Println(has)
}
