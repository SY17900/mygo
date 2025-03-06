package course

// var prereqs = map[string][]string{
// 	"algorithms": {"data structures"},
// 	"calculus":   {"linear algebra"},
// 	"compilers": {
// 		"data structures",
// 		"formal languages",
// 		"computer organization",
// 	},
// 	"data structures":       {"discrete math"},
// 	"databases":             {"data structures"},
// 	"discrete math":         {"intro to programming"},
// 	"formal languages":      {"discrete math"},
// 	"networks":              {"operating systems"},
// 	"operating systems":     {"data structures", "computer organization"},
// 	"programming languages": {"data structures", "computer organization"},
// }

func GetOrder(mp map[string][]string) (res []string) {
	visited := map[string]struct{}{}
	var dfs func([]string)

	dfs = func(courses []string) {
		for _, course := range courses {
			if _, ok := visited[course]; !ok {
				visited[course] = struct{}{}
				dfs(mp[course])
				res = append(res, course)
			}
		}
	}

	var courses []string
	for course := range mp {
		courses = append(courses, course)
	}
	dfs(courses)

	return
}
