package exam

import "fmt"

func Grid() {
	var m, n, ans int
	fmt.Scan(&m, &n)
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for i := range grid {
		for j := range grid[0] {
			fmt.Scan(&grid[i][j])
		}
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int)
	dfs = func(i, j int) {
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		if i+1 < m && grid[i+1][j] == 1 {
			dfs(i+1, j)
		}
		if j+1 < n && grid[i][j+1] == 1 {
			dfs(i, j+1)
		}
		if i-1 >= 0 && grid[i-1][j] == 1 {
			dfs(i-1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == 1 {
			dfs(i, j-1)
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if visited[i][j] {
				continue
			}
			if grid[i][j] == 0 {
				visited[i][j] = true
				continue
			}
			// grid[i][j] == 1
			ans++
			dfs(i, j)
		}
	}

	fmt.Println(ans)
}
