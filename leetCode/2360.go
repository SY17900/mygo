package leetcode

func LongestCycle(edges []int) int {
	ans, curTime := -1, 1
	visit := make([]int, len(edges))
	for start := range edges {
		startTime := curTime
		curNode := start
		for curNode != -1 && visit[curNode] == 0 {
			visit[curNode] = curTime
			curTime++
			curNode = edges[curNode]
		}
		if curNode == -1 {
			continue
		}
		if visit[curNode] >= startTime {
			ans = max(ans, curTime-visit[curNode])
		}
	}

	return ans
}
