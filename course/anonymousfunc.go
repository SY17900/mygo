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

func StudyPath(mp map[string][]string) []string {
	var order []string
	seen := map[string]bool{}
	var dfs func(items []string)

	dfs = func(items []string) {
		for _, item := range items {
			// 找到该课程的所有依赖，然后再填入该课程
			if !seen[item] {
				seen[item] = true
				dfs(mp[item])
				order = append(order, item)
			}
		}
	}

	var courses []string
	for course := range mp {
		courses = append(courses, course)
	}

	dfs(courses)

	return order
}
