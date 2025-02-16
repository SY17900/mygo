package graphs

import (
	"fmt"
)

var ans [][]int
var path []int

func dfs(graph [][]int, cur, target int) {
	if cur == target {
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		return
	}

	for i := 1; i <= target; i++ {
		if graph[cur][i] == 1 {
			path = append(path, i)
			dfs(graph, i, target)
			path = path[:len(path)-1]
		}
	}
}

func AllPaths() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		var s, t int
		fmt.Scanf("%d %d", &s, &t)
		graph[s][t] = 1
	}

	path = append(path, 1)
	dfs(graph, 1, n)

	if len(ans) == 0 {
		fmt.Println(-1)
	} else {
		for _, p := range ans {
			for i := 0; i < len(p)-1; i++ {
				fmt.Print(p[i], " ")
			}
			fmt.Println(p[len(p)-1])
		}
	}
}
