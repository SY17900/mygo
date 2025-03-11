package backtracking

import "sort"

func FindItinerary(tickets [][]string) (ans []string) {
	type dest struct {
		city string
		used bool
	}
	mp := make(map[string][]dest)
	for _, ticket := range tickets {
		if _, ok := mp[ticket[0]]; ok {
			mp[ticket[0]] = append(mp[ticket[0]], dest{city: ticket[1]})
			continue
		}
		mp[ticket[0]] = []dest{{city: ticket[1]}}
	}
	for key := range mp {
		sort.Slice(mp[key], func(i int, j int) bool {
			return mp[key][i].city < mp[key][j].city
		})
	}

	var dfs func(string) bool
	rec := []string{"JFK"}
	dfs = func(start string) bool {
		if len(rec) == len(tickets)+1 {
			ans = rec
			return true
		}
		for i, destination := range mp[start] {
			if destination.used {
				continue
			}
			rec = append(rec, destination.city)
			mp[start][i].used = true
			if dfs(destination.city) {
				return true
			}
			mp[start][i].used = false
			rec = rec[:len(rec)-1]
		}

		return false
	}

	dfs("JFK")
	return
}

func NQueens(n int) (ans [][]string) {
	rec := make([][]bool, n)
	for row := range rec {
		rec[row] = make([]bool, n)
	}
	ifvalid := func(row, col int) bool {
		for i := range row {
			if rec[i][col] {
				return false
			}
		}
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if rec[i][j] {
				return false
			}
		}
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if rec[i][j] {
				return false
			}
		}

		return true
	}
	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			board := []string{}
			for _, i := range rec {
				temp := []byte{}
				for _, j := range i {
					if !j {
						temp = append(temp, '.')
					} else {
						temp = append(temp, 'Q')
					}
				}
				board = append(board, string(temp))
			}
			ans = append(ans, board)
			return
		}
		for i := range n {
			if !ifvalid(row, i) {
				continue
			}
			rec[row][i] = true
			dfs(row + 1)
			rec[row][i] = false
		}
	}

	dfs(0)
	return
}
