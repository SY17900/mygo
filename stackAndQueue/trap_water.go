package stackandqueue

import (
	"container/heap"
	"fmt"
)

// 二维接雨水
func TrapWater2D(height []int) int {
	stack := []int{}
	ans := 0

	for position, n := range height {
		for len(stack) > 0 && n > height[stack[len(stack)-1]] {
			button := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				h := min(height[stack[len(stack)-1]], n) - button
				w := position - stack[len(stack)-1] - 1
				fmt.Println(h, w, h*w)
				ans += h * w
			}
		}
		if len(stack) != 0 && n == height[stack[len(stack)-1]] {
			stack[len(stack)-1] = position
		} else {
			stack = append(stack, position)
		}
	}

	return ans
}

// 三维接雨水
type cell struct {
	height, x, y int
}
type hp []cell

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Less(i, j int) bool {
	return (*h)[i].height < (*h)[j].height
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(cell))
}

func (h *hp) Pop() any {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

var dir4 = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func TrapWater3D(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	h := hp{}
	ans := 0
	for i, row := range heightMap {
		for j, height := range heightMap[i] {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				h = append(h, cell{height: height, x: i, y: j})
				row[j] = -1
			}
		}
	}
	heap.Init(&h)

	for len(h) > 0 {
		lowest := heap.Pop(&h).(cell)
		for _, dir := range dir4 {
			x, y := lowest.x+dir.x, lowest.y+dir.y
			if 0 <= x && x <= m && 0 <= y && y <= n && heightMap[x][y] != -1 {
				if heightMap[x][y] < lowest.height {
					ans += lowest.height - heightMap[x][y]
				}
				heap.Push(&h, cell{max(heightMap[x][y], lowest.height), x, y})
				heightMap[x][y] = -1
			}
		}
	}
	return ans
}
