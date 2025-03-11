package backtracking

import "sort"

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

func CombineSum(n, k int) (ans [][]int) {
	var dfs func(int)
	rec := make([]int, 0, k)
	rest := n
	dfs = func(i int) {
		if len(rec) > k {
			return
		}
		if len(rec) == k {
			if rest == 0 {
				temp := make([]int, k)
				copy(temp, rec)
				ans = append(ans, temp)
				return
			} else if rest > 0 {
				return
			}
		}
		for next := i; next <= 9 && rest-next >= 0; next++ {
			rec = append(rec, next)
			rest -= next
			dfs(next + 1)
			rest += next
			rec = rec[:len(rec)-1]
		}
	}

	dfs(1)
	return
}

func Letters(digits string) (ans []string) {
	mp := []string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	var dfs func(int)
	rec := make([]byte, 0, len(digits))
	dfs = func(i int) {
		if i == len(digits) {
			ans = append(ans, string(rec))
			return
		}
		for _, letter := range mp[(int)(digits[i]-'0')] {
			rec = append(rec, byte(letter))
			dfs(i + 1)
			rec = rec[:len(rec)-1]
		}
	}

	dfs(0)
	return
}

func CombineSumD(candidates []int, target int) (ans [][]int) {
	var dfs func(int)
	rec := make([]int, 0)
	dfs = func(i int) {
		if target == 0 {
			temp := make([]int, len(rec))
			copy(temp, rec)
			ans = append(ans, temp)
			return
		}
		for next := i; next < len(candidates); next++ {
			if candidates[next] > target {
				break
			}
			rec = append(rec, candidates[next])
			target -= candidates[next]
			dfs(next)
			target += candidates[next]
			rec = rec[:len(rec)-1]
		}
	}

	dfs(0)
	return
}

func CombineSumND(candidates []int, target int) (ans [][]int) {
	var dfs func(int)
	rec := make([]int, 0, len(candidates))
	dfs = func(i int) {
		if target == 0 {
			temp := make([]int, len(rec))
			copy(temp, rec)
			ans = append(ans, temp)
			return
		}
		for next := i; next < len(candidates); next++ {
			if candidates[next] > target {
				break
			}
			if next > i && candidates[next] == candidates[next-1] {
				continue
			}
			rec = append(rec, candidates[next])
			target -= candidates[next]
			dfs(next + 1)
			target += candidates[next]
			rec = rec[:len(rec)-1]
		}
	}

	sort.Ints(candidates)
	dfs(0)
	return
}

func MakePalindrome(s string) (ans [][]string) {
	var dfs func(int)
	var ifPalindrome func(int, int) bool
	rec := make([]string, 0)
	ifPalindrome = func(i, j int) bool {
		if i == j {
			return true
		}
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	dfs = func(i int) {
		if i == len(s) {
			temp := make([]string, len(rec))
			copy(temp, rec)
			ans = append(ans, temp)
			return
		}
		for next := i; next < len(s); next++ {
			if !ifPalindrome(i, next) {
				continue
			}
			rec = append(rec, string(s[i:next+1]))
			dfs(next + 1)
			rec = rec[:len(rec)-1]
		}
	}

	dfs(0)
	return
}
