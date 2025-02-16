package dp

func Backpack01(costs, values []int, vol int) int {
	rec := make([]int, vol+1)
	cnt := len(costs)
	for i := 0; i < cnt; i++ {
		for j := vol; j >= costs[i]; j-- {
			rec[j] = max(rec[j], rec[j-costs[i]]+values[i])
		}
	}

	return rec[vol]
}
