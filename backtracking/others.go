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
